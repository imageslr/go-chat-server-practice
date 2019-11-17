/**
CommandReader 将命令包装为字节码并写入到 io.Writer 中，这个 io.Writer 实际上是 net.Conn
*/

package protocol

import (
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
	// TODO
	return nil
}
