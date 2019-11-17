package server

// 定义一个 interface，让 server 的行为更加清晰明了
type ChatServer interface {
	Listen(address string) error         // 监听客户端发来的连接
	Broadcast(command interface{}) error // 将客户端发来的消息广播给所有客户端
	Start()                              // 为客户端提供服务
	Close()                              // 关闭服务器
}
