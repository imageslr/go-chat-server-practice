package server

import (
	"errors"
	"go-chat-server-practice/protocol"
	"io"
	"log"
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
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s.listener = l

	log.Printf("Listening on %v", address)

	return nil
}

// 取消监听
func (s *TcpChatServer) Close() {
	_ = s.listener.Close()
}

// 不停的接受客户端的连接请求，并为其提供服务
func (s *TcpChatServer) Start() {
	for {
		// 调用 s.listener.Accept() 得到一个与客户端的连接 conn
		// 调用 s.accept(conn) 将该连接转为 client 结构体，保存在 s.clients 中
		// 调用 go s.serve(client) 在一个新的协程中为该客户端提供服务
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		client := s.accept(conn)
		go s.serve(client)
	}
}

// 将客户端发来的消息广播给所有客户端，command 参数其实是 MessageCommand，这里是为了实现接口定义
func (s *TcpChatServer) Broadcast(command interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, client := range s.clients {
		err := client.writer.Write(command)
		if err != nil {
			log.Printf("Broadcast to client %v failed: %v", client.name, err)
		}
	}

	return nil
}

// 将与客户端的连接转为 client 结构体，保存在 s.clients 中
func (s *TcpChatServer) accept(conn net.Conn) *client {
	log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.clients)+1)

	// 这里一定要加锁，因为可能有多个协程访问 clients（如 `go s.Broadcast`、`go s.serve`里的 `s.remove`）
	s.mutex.Lock()
	defer s.mutex.Unlock()

	client := &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	}
	s.clients = append(s.clients, client)

	return client
}

// 为客户端提供服务：读取客户端发来的命令，执行相应的操作；在客户端断开连接时，移除该连接
// 客户端是否断开连接可以通过 Read() 返回 EOF 错误来判断
// 具体实现和 protocol.CommandReader.ReadAll() 很相似
func (s *TcpChatServer) serve(client *client) {
	r := protocol.NewCommandReader(client.conn)

	for {
		command, err := r.Read()
		if err == io.EOF {
			s.remove(client) // io.EOF 表示客户端主动断开连接，这时移除该连接，结束服务
			break
		} else if err != nil {
			log.Printf("Read error: %v", err) // 其他情况不要断开连接，比如：客户端发送的命令格式错误
			continue
		}

		switch v := command.(type) {
		case protocol.SendCommand:
			// 在一个协程里异步广播消息，不要阻塞当前客户端的下一条命令
			go s.Broadcast(protocol.MessageCommand{
				Message: v.Message,
				Name:    client.name,
			})
		case protocol.NameCommand:
			client.name = v.Name
		}
	}
}

// 断开与某个客户端的连接
func (s *TcpChatServer) remove(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, v := range s.clients {
		if v == client {
			// 这里注意会直接修改 for 循环中 range 使用的切片
			// 下一步里：i 递增 1，v 会在新的切片中使用下标 i 访问
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}

	log.Printf("Closing connection from %v", client.conn.RemoteAddr().String())
	_ = client.conn.Close()
}
