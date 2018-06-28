package wlog

import (
	"fmt"
	"sync"
)

type ChannelFactory struct {
	mu       *sync.Mutex
	builders map[string]ChannelBuilderInterface
}

var g_factory *ChannelFactory
func init() {
	g_factory = &ChannelFactory{new(sync.Mutex), make(map[string]ChannelBuilderInterface, 10)}	
	//buildin
	AddBuilder(STDOUT, &StdoutChannelBuilder{})
	AddBuilder(MSTDOUT, &MStdoutChannelBuilder{})
	AddBuilder(MSFC, &MaxSizeFileChannelBuilder{})
}

func AddBuilder(name string, b ChannelBuilderInterface) {
	g_factory.mu.Lock()
	defer g_factory.mu.Unlock()
	_,ok := g_factory.builders[name]
	if ok {
		return
	}
	g_factory.builders[name] = b
}

func AddBuilderForce(name string, b ChannelBuilderInterface) {
	g_factory.mu.Lock()
	defer g_factory.mu.Unlock()
	g_factory.builders[name] = b
}

func GetBuilder(cc *ChannelConfig) (ChannelBuilderInterface, bool) {
	g_factory.mu.Lock()
	defer g_factory.mu.Unlock()
	b, ok := g_factory.builders[cc.Ctype]
	if ok {
		return b, true
	}	
	fmt.Println("get channel builder", cc.Ctype, "error")
	return nil, false
}

func GetChannel(cc *ChannelConfig) Channel {
	b, ok := GetBuilder(cc)
	if !ok {
		return nil
	}
	c := b.CreateChannel(cc)
	return c
}