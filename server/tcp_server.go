package server

import (
	"errors"
	"go-chat-server-practice/protocol"
	"net"
	"sync"
)

var (
	UnknownClient = errors.New("Unknown client")
)

// client 表示一个客户端连接，包括 tcp 连接、客户端名称与向客户端写入命令的 writer
type client struct {
	conn   net.Conn
	name   string
	writer *protocol.CommandWriter
}

/**
TcpChatServer 接受客户端的连接请求，并保存所有连接（上面定义的 client 结构体）。

TcpChatServer 的工作流程为：
1. Listen：启动一个服务器，监听一个端口
2. Accept & Serve：与客户端建立一个连接，并为其提供服务
3. Remove：在客户端退出连接后，删除该客户的的连接
4. Close：停止监听端口，关闭服务器

TcpChatServer 与客户端的交互有：
1. 接受客户端发来的消息，并广播给其他客户端
2. 设置某个客户端的名称
*/
type TcpChatServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex // 所有对 clients 的操作都要加锁，防止竞态情况
}

func NewServer() *TcpChatServer {
	return &TcpChatServer{
		mutex: &sync.Mutex{},
	}
}

// 监听一个端口
func (s *TcpChatServer) Listen(address string) error {
	// TODO
	return nil
}

// 取消监听
func (s *TcpChatServer) Close() {
	// TODO
}

// 不停的接受客户端的连接请求，并为其提供服务
func (s *TcpChatServer) Start() {
	// TODO
	// 调用 s.listener.Accept() 得到一个与客户端的连接 conn
	// 调用 s.accept(conn) 将该连接转为 client 结构体，保存在 s.clients 中
	// 调用 go s.serve(client) 在一个新的协程中为该客户端提供服务
}

// 将客户端发来的消息广播给所有客户端，command 参数的类型其实是 MessageCommand
func (s *TcpChatServer) Broadcast(command interface{}) error {
	// TODO
	return nil
}

// 将与客户端的连接转为 client 结构体，保存在 s.clients 中
func (s *TcpChatServer) accept(conn net.Conn) *client {
	// TODO
	return nil
}

// 为客户端提供服务：读取客户端发来的命令，执行相应的操作；在客户端断开连接时，移除该连接
// 客户端是否断开连接可以通过 Read() 返回 EOF 错误来判断
func (s *TcpChatServer) serve(client *client) {
	// TODO
}

// 断开与某个客户端的连接
func (s *TcpChatServer) remove(client *client) {
	// TODO
}
