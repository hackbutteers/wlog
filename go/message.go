package wlog

import (
	"bytes"
	"time"
)

type MessageLite struct {

	LogName       *string

	Time          time.Time

	LogLevel       Level

}

type Message struct {

	LogName       *string

	Time          time.Time

	LogLevel       Level

	Raw           *bytes.Buffer
	Formated      *bytes.Buffer
}

func NewMessageLite(s *string, t time.Time, l Level ) *MessageLite {
	return &MessageLite{s, t, l}
}

func NewMessage(m *MessageLite) *Message {
	return &Message{ m.LogName, m.Time, m.LogLevel, new(bytes.Buffer), new(bytes.Buffer)}
}
