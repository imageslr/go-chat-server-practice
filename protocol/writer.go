/**
CommandReader 将命令包装为字节码并写入到 io.Writer 中，这个 io.Writer 实际上是 net.Conn
*/

package protocol

import (
	"encoding/gob"
	"io"
)

type CommandWriter struct {
	writer *gob.Encoder
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: gob.NewEncoder(writer),
	}
}

/**
Write 向 w.writer 中写入一条命令
command 为 protocol/command 中的 XXXCommand 结构体，但最终写入的命令应该是符合格式约定的字符串
*/
func (w *CommandWriter) Write(command interface{}) error {
	var err error

	switch v := command.(type) {
	case SendCommand:
		err = w.writer.Encode(Command{Type: "SEND", Message: v.Message})
	case MessageCommand:
		err = w.writer.Encode(Command{Type: "MESSAGE", Name: v.Name, Message: v.Message})
	case NameCommand:
		err = w.writer.Encode(Command{Type: "NAME", Name: v.Name})
	default:
		err = UnknownCommand
	}

	return err
}
