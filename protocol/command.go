/**
客户端与服务端通过 TCP 传输的是字符串，因此需要规定一个将字符串解析为命令的协议。

一条命令的格式如下所示：
	[命令类型] [参数1] [参数2] ... [参数n]\n
每条命令以换行符结尾。

命令一共有三种：
* NAME：客户端设置用户名
* SEND：客户端发送聊天消息
* MESSAGE：服务端广播聊天消息给其他用户

比如客户端发送聊天消息的命令为：
	SEND somemessage\n
服务端广播消息给其他用户的命令为：
	MESSAGE username somemessage\n

以下定义的结构体用来标识命令的类型，以及相关的参数。
*/

package protocol

import "errors"

var (
	UnknownCommand = errors.New("Unknown command")
)

// 客户端向服务端上发的命令
type SendCommand struct {
	Message string
}

// 客户端向服务端上发的命令
type NameCommand struct {
	Name string
}

// 服务端向客户端下发的命令
type MessageCommand struct {
	Name    string
	Message string
}
