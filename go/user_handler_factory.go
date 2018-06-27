package wlog

import (
	"sync"
)

const (
	DUH = "default_text_handler"
)
type UserHandlerBuilderInterface interface {
	CreateUH(conf *LogChannelConfig) UserHandler
}

type UserHandlerFactory struct {
	mu       *sync.Mutex
	builders map[string]UserHandlerBuilderInterface
}

var g_uh_factory *UserHandlerFactory
func init() {
	g_uh_factory = &UserHandlerFactory{new(sync.Mutex), make(map[string]UserHandlerBuilderInterface, 10)}	
	//buildin
	AddUHBuilder(DUH, &DefaultTextUserHandlerBuilder{})
}

func AddUHBuilder(name string, b UserHandlerBuilderInterface) {
	g_uh_factory.mu.Lock()
	defer g_uh_factory.mu.Unlock()
	_,ok := g_uh_factory.builders[name]
	if ok {
		return
	}
	g_uh_factory.builders[name] = b
}

func AddUHBuilderForce(name string, b UserHandlerBuilderInterface) {
	g_uh_factory.mu.Lock()
	defer g_uh_factory.mu.Unlock()
	g_uh_factory.builders[name] = b
}

func GetUHBuilder(name string) (UserHandlerBuilderInterface, bool) {
	g_uh_factory.mu.Lock()
	defer g_uh_factory.mu.Unlock()
	b, ok := g_uh_factory.builders[name]
	if ok {
		return b, true
	}	
	return nil, false
}

func GetUH(cc *LogChannelConfig) UserHandler {
	b, ok := GetUHBuilder(cc.Handlername)
	if !ok {
		return nil
	}
	c := b.CreateUH(cc)
	return c
}