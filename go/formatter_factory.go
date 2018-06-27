package wlog

import (
	"sync"
)


type FormatterBuilderInterface interface {
	CreateFormatter(lc *LogChannelConfig)Formatter
}


type FormatterFactory struct {
	mu       *sync.Mutex
	builders map[string]FormatterBuilderInterface
}

var g_formatter_factory *FormatterFactory


func init() {
	g_formatter_factory = &FormatterFactory{new(sync.Mutex), make(map[string]FormatterBuilderInterface, 10)}	
	//buildin
	AddFormatterBuilder(DTESTF, &DefaultTestFormatterBuilder{})
}

func AddFormatterBuilder(name string, b FormatterBuilderInterface) {
	g_formatter_factory.mu.Lock()
	defer g_formatter_factory.mu.Unlock()
	_,ok := g_formatter_factory.builders[name]
	if ok {
		return
	}
	g_formatter_factory.builders[name] = b
}

func AddFormatterBuilderForce(name string, b FormatterBuilderInterface) {
	g_formatter_factory.mu.Lock()
	defer g_formatter_factory.mu.Unlock()
	g_formatter_factory.builders[name] = b
}

func GetFormatterBuilder(name string) (FormatterBuilderInterface, bool) {
	g_formatter_factory.mu.Lock()
	defer g_formatter_factory.mu.Unlock()
	b, ok := g_formatter_factory.builders[name]
	if ok {
		return b, true
	}	
	return nil, false
}

func GetFormatter(cc *LogChannelConfig) Formatter {
	b, ok := GetFormatterBuilder(cc.Formatter)
	if !ok {
		return nil
	}
	c := b.CreateFormatter(cc)
	return c
}