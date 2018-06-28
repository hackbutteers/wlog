package wlog

import (
	"fmt"
	"bytes"
)

type DefaultTextHandler struct {
	Name  string
}

func NewDefaultTextHandler() *DefaultTextHandler {
	return &DefaultTextHandler{"default_test_handler"}
}
func (t *DefaultTextHandler)Formatf(buffer *bytes.Buffer, format string, v...interface{}) {
	fmt.Fprintf(buffer, format, v...)
}

func (t *DefaultTextHandler)Format(buffer *bytes.Buffer, v...interface{}) {
	fmt.Fprint(buffer, v...)
}

func (t *DefaultTextHandler)getName() string{
	return t.Name
}

type DefaultTextUserHandlerBuilder struct {}

func (d*DefaultTextUserHandlerBuilder)CreateUH(conf *LogChannelConfig) UserHandler {
	return NewDefaultTextHandler()
}