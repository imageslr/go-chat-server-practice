/**
CommandReader 从 io.Reader 中读取并解析命令，这个 io.Reader 实际上是 net.Conn
*/

package protocol

import (
	"encoding/gob"
	"io"
)

type CommandReader struct {
	reader *gob.Decoder
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: gob.NewDecoder(reader),
	}
}

/**
Read 从 r.reader 中读取一条命令
interface{} 是 protocol/command 中定义的 XXXCommand 结构体
*/
func (r *CommandReader) Read() (interface{}, error) {
	command := Command{}
	if err := r.reader.Decode(&command); err != nil {
		return nil, err
	}

	switch command.Type {
	case "SEND":
		return SendCommand{Message: command.Message}, nil
	case "NAME":
		return NameCommand{Name: command.Name}, nil
	case "MESSAGE":
		return MessageCommand{Message: command.Message, Name: command.Name}, nil
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
