/**
CommandReader 从 io.Reader 中读取并解析命令，这个 io.Reader 实际上是 net.Conn
*/

package protocol

import (
	"bufio"
	"io"
)

type CommandReader struct {
	reader *bufio.Reader
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: bufio.NewReader(reader),
	}
}

/**
Read 从 r.reader 中读取一条命令
interface{} 是 protocol/command 中定义的 XXXCommand 结构体
*/
func (r *CommandReader) Read() (interface{}, error) {
	// TODO
	return nil, UnknownCommand
}

/**
ReadAll 从 r.reader 中读取所有命令。interface{} 同上
*/
func (r *CommandReader) ReadAll() ([]interface{}, error) {
	// TODO
	return nil, nil
}
