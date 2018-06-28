package wlog

import (
	"bytes"
)

type UserHandler interface {
	Formatf(buffer *bytes.Buffer, format string, v...interface{})
	Format(buffer *bytes.Buffer, v...interface{})
	getName() string
}

type Formatter interface {
	Format(m *Message)
	CompileConf(str string) bool
}


type FormatterBuilderInterface interface {
	CreateFormatter(lc *LogChannelConfig)Formatter
}



type UserHandlerBuilderInterface interface {
	CreateUH(conf *LogChannelConfig) UserHandler
}

type Channel interface {
	
	Log(m *Message)

	Flush()

	SetProperty(k string, v string)
	
	GetProperty(k string) string
	
	GetAllProperty() map[string]string
	
	AddChannel(c Channel)
	
	GetAllChannel() []Channel
	
	GetChannel(name string) Channel
}


type ChannelBuilderInterface interface {
	CreateChannel(conf *ChannelConfig) Channel
}
