package main

import (
	"time"
	"tunnel/client"
	"tunnel/server"
)

func main() {
	go server.RunServer(7001)
	time.Sleep(1 * time.Second)

	go client.RunClient(7001, 7002, "127.0.0.1:7003", "1")
	client.RunClient(7001, 7005, "127.0.0.1:7006", "2")
}
