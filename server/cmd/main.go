/**
在此目录下运行 go run main.go 启动服务器
*/

package main

import (
	"go-chat-server-practice/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")

	// start the server
	s.Start()
}
