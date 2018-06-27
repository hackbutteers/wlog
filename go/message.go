package wlog

import (
	"os"
	"bytes"
	"time"
)

type MessageLite struct {

	LogName       *string

	Time          time.Time

	LogLevel       Level

	Pid            int

}

type Message struct {

	LogName       *string

	Time          time.Time

	LogLevel       Level

	Pid            int

	Raw           *bytes.Buffer
	Formated      *bytes.Buffer
}

func NewMessageLite(s *string, t time.Time, l Level ) *MessageLite {
	return &MessageLite{s, t, l,os.Getegid()}
}

func NewMessage(m *MessageLite, raw *bytes.Buffer, formated *bytes.Buffer) *Message {
	return &Message{ m.LogName, m.Time, m.LogLevel, m.Pid, raw, formated}
}
