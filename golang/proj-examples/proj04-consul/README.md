docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0