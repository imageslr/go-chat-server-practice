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
	command, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}

	command = command[:len(command)-1] // bufio.ReadString 方法返回的结果包含界定符，去掉

	switch command {
	case "SEND":
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return SendCommand{Message: message[:len(message)-1]}, nil
	case "NAME":
		name, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return NameCommand{Name: name[:len(name)-1]}, nil
	case "MESSAGE":
		name, err := r.reader.ReadString(' ')
		if err != nil {
			return nil, err
		}
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return MessageCommand{Message: message[:len(message)-1], Name: name[:len(name)-1]}, nil
	}

	return nil, UnknownCommand
}

/**
ReadAll 从 r.reader 中读取所有命令。interface{} 同上
*/
func (r *CommandReader) ReadAll() ([]interface{}, error) {
	commands := []interface{}{}

	for {
		command, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return commands, err
		}

		commands = append(commands, command)
	}

	return commands, nil
}
