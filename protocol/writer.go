/**
CommandReader 将命令包装为字节码并写入到 io.Writer 中，这个 io.Writer 实际上是 net.Conn
*/

package protocol

import (
	"fmt"
	"io"
)

type CommandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: writer,
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
		_, err = w.writer.Write([]byte(fmt.Sprintf("SEND %s\n", v.Message)))
	case MessageCommand:
		_, err = w.writer.Write([]byte(fmt.Sprintf("MESSAGE %s %s\n", v.Name, v.Message)))
	case NameCommand:
		_, err = w.writer.Write([]byte(fmt.Sprintf("NAME %s\n", v.Name)))
	default:
		err = UnknownCommand
	}

	return err
}
