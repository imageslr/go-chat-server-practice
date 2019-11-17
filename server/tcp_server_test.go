package server_test

import (
	"go-chat-server-practice/protocol"
	"net"
	"testing"
)

/**
这个代码有问题
测试服务器能否正确发送广播
但是客户端 1 发送消息后，客户端再 read 的时候会阻塞，暂时不知道怎么解决
*/
func TestServer(t *testing.T) {
	tests := []struct {
		name    string
		message string
		result  interface{}
	}{
		{
			"user1",
			"hahaha",
			protocol.MessageCommand{"user1", "hahaha"},
		},
	}

	// 需要自己先启动服务器 `go run server/cmd/main.go`
	//go func() {
	//	s := server.NewServer()
	//	err := s.Listen(":3333")
	//	if err != nil {
	//		t.Errorf("Ubable to start the server: %v", err)
	//	}
	//	defer s.Close()
	//
	//	s.Start()
	//}()

	// 启动客户端 1
	conn1, err := net.Dial("tcp", ":3333")
	if err != nil {
		t.Errorf("Client 1 ubable to connect to the server: %v", err)
	}
	defer conn1.Close()

	// 启动客户端 2
	conn2, err := net.Dial("tcp", ":3333")
	if err != nil {
		t.Errorf("Client 2 unable to connect to the server: %v", err)
	}
	defer conn2.Close()

	for _, test := range tests {
		r := protocol.NewCommandReader(conn2)
		w := protocol.NewCommandWriter(conn1)

		// 客户端 1 发送一条消息
		command := protocol.MessageCommand{Name: test.name, Message: test.message}
		if err := w.Write(command); err != nil {
			t.Errorf("Ubable to send message: %v", err)
		}

		t.Log("Send message success")

		// 客户端 2 应该收到对应的命令
		result, err := r.Read()
		if err != nil {
			t.Errorf("Unable to read command, error %v", err)
		} else if result != test.result {
			t.Errorf("Command output is not the same: %v %v", result, test.result)
		}
	}
}
