package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/kubectl/pkg/scheme"
	client "k8s/common"
	ws "k8s/web-ssh/stream"
	"net/http"
)

// ssh流式处理器
type streamHandler struct {
	// websocket连接
	wsConn *ws.WebsocketConn
	// 设置终端的大小
	resizeEvent chan remotecommand.TerminalSize
}

// web终端发来的包信息
type xtermMessage struct {
	// 类型 -> resize客户端调整终端的input输入
	MsgType string `json:"type"`
	// msgtype=input 情况下使用
	Input string `json:"input"`
	// msgtype=resize 情况下使用
	Rows uint16 `json:"rows"`
	// msgtype=resize 情况下使用
	Cols uint16 `json:"cols"`
}

// executor回调来获取web是否resize
func (handler *streamHandler) Next() (size *remotecommand.TerminalSize) {
	ret := <-handler.resizeEvent
	size = &ret
	return size
}

// executor 读取web端的输入信息
func (handler *streamHandler) Read(r []byte) (size int, err error) {
	var (
		msg     *ws.WebSocketMessage
		xterMsg *xtermMessage
	)
	// 读取来自web端发来的输入信息
	if msg, err = handler.wsConn.WsRead(); err != nil {
		return 0, err
	}
	// 解析客户端的请求信息
	if err = json.Unmarshal(msg.Date, xterMsg); err != nil {
		return 0, err
	}
	// web-ssh 调整终端的大小
	if xterMsg.MsgType == "resize" {
		// 传值到remote的terminalsize中然后存放在channel
		handler.resizeEvent <- remotecommand.TerminalSize{
			Width:  xterMsg.Cols,
			Height: xterMsg.Rows,
		}
	} else if xterMsg.MsgType == "input" {
		// web-ssh 终端输入了字符 将它拷贝到数组中
		size = len(xterMsg.Input)
		copy(r, xterMsg.Input)
	}
	return
}

// executor 写入web输出端
func (handle *streamHandler) Write(r []byte) (size int, err error) {
	var (
		copyData []byte
	)
	// 生成副本
	copyData = make([]byte, len(r))
	// 拷贝数据
	copy(copyData, r)
	// 计算大小
	size = len(r)
	// 执行发送
	err = handle.wsConn.WsSend(websocket.TextMessage, copyData)
	return
}

// 处理container输出的handler
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		wsConn   *ws.WebsocketConn
		handler  *streamHandler
		executor remotecommand.Executor
	)

	fmt.Println("有request进来了")

	// 解析Get请求
	if err = r.ParseForm(); err != nil {
		return
	}

	// 获取需要的所有容器的信息
	podNS := r.Form.Get("podns")
	podName := r.Form.Get("podname")
	containerName := r.Form.Get("containername")

	// 获取websocket的长连接
	if wsConn, err = ws.InitWebsocket(w, r); err != nil {
		return
	}

	podNS = "default"
	podName = "busybox"
	containerName = "busybox"

	// 获取k8s restclient信息
	res, _ := client.NewResAndLog()

	// 获取k8s rest信息
	rest := client.GetRestconf()

	// 获取podclient的restclient
	sshRequest := res.CoreV1().RESTClient().Post().
		Resource("pods").Name(podName).
		Namespace(podNS).SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			TypeMeta:  metav1.TypeMeta{},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
			Container: containerName,
			Command:   []string{"sh"},
		}, scheme.ParameterCodec)

	// 创建到容器的连接
	if executor, err = remotecommand.NewSPDYExecutor(rest, "POST", sshRequest.URL()); err != nil {
		goto ERR
	}

	// 配置和容器之间的数据流处理和回调
	handler = &streamHandler{
		wsConn:      wsConn,
		resizeEvent: make(chan remotecommand.TerminalSize),
	}

	// 启动shell的流传输（将输出的方式全部使用我们自定义的handle处理方式）
	if err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		Tty:               true,
		TerminalSizeQueue: handler,
	}); err != nil {
		goto ERR
	}

	return

ERR:
	fmt.Println(err)
	wsConn.WsClose()
}

func htmlhandle(w http.ResponseWriter, r *http.Request) {
	T := template.New("index")
	data, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Println(err)
	}
	T.Parse(string(data))

	data1 := map[string]string{
		"name": "测试",
	}
	T.Execute(w, data1)
}

func main() {
	fmt.Println("开始监听7777端口")
	// 创建监听的端口号
	http.HandleFunc("/ssh", wsHandler)
	http.HandleFunc("/", htmlhandle)
	http.ListenAndServe(":7777", nil)
}
