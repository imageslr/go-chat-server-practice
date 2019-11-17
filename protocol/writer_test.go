package protocol_test

import (
	"bytes"
	"go-chat-server-practice/protocol"
	"testing"
)

func TestCommandWriter(t *testing.T) {
	tests := []struct {
		input   interface{}
		results string
	}{
		{
			protocol.SendCommand{"test"},
			"SEND test\n",
		},
		{
			protocol.MessageCommand{"user1", "hello"},
			"MESSAGE user1 hello\n",
		},
	}

	buf := new(bytes.Buffer)
	for _, test := range tests {
		buf.Reset()
		writer := protocol.NewCommandWriter(buf)
		err := writer.Write(test.input)

		if err != nil {
			t.Errorf("Unable to write command, error %v", err)
		} else if test.results != buf.String() {
			t.Errorf("Command output is not the same: %v %v", buf.String(), test.results)
		}
	}
}
