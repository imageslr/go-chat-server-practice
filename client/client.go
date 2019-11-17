package client

import "go-chat-server-practice/protocol"

type messageHandler func(string)

// 定义一个 interface，让 client 的行为更加清晰明了
type ChatClient interface {
	Dial(address string) error      // 连接到服务端
	Start()                         // 获取服务端发来的命令，执行相应的操作
	Close()                         // 关闭连接
	Send(command interface{}) error // 向服务端发送一条命令
	SetName(name string) error
	SendMessage(message string) error
	Incoming() chan protocol.MessageCommand // 服务端广播的消息会发送到这个 channel
}
