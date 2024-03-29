package main

import (
	"log"
	"runtime"
)

var intMap map[int]int
var cnt = 8192

func main() {
	printMemStats()

	initData()
	runtime.GC()
	printMemStats()

	log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap)
	}
	log.Println(len(intMap))
}

func initData() {
	intMap = make(map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v, TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1014, m.Sys/1024, m.NumGC)
}
