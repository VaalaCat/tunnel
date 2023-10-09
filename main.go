package main

import (
	"time"

	"github.com/VaalaCat/tunnel/client"
	"github.com/VaalaCat/tunnel/server"
	"github.com/VaalaCat/tunnel/testserver"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	go server.RunServer(7001)
	go testserver.RunGinServer(7003)
	go testserver.RunGinServer(7006)
	time.Sleep(3 * time.Second)

	go client.RunClient("localhost", 7001, 7002, "127.0.0.1:7003", "1")
	time.Sleep(1 * time.Second)
	client.DeleteClient("localhost", 7001, "1")
	client.RunClient("localhost", 7001, 7005, "127.0.0.1:7006", "2")
}
