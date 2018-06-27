package wlog

import (
	"fmt"
	"sync"
)

type StdoutChannel struct {
	name string
}

type MStdoutChannel struct {
	mu sync.Mutex
	name string
}

func (sc *StdoutChannel)Log(m *Message) {
	b := m.Formated.String()
	fmt.Print(b)
}

func (sc *StdoutChannel)Flush() {
}

func (sc *StdoutChannel)SetProperty(k string, v string) {
	if k == "name" {
		sc.name = v
	}
}
	
func (sc *StdoutChannel)GetProperty(k string) string {
	if k == "name" {
		return sc.name
	}
	return ""
}

func (sc *StdoutChannel)GetAllProperty() map[string]string {
	return map[string]string{}
}

func (sc *StdoutChannel)AddChannel(c Channel) {

}

func (sc *StdoutChannel)GetAllChannel() []Channel {
	return nil
}

func (sc *StdoutChannel)GetChannel(name string) Channel {
	return nil
}

type StdoutChannelBuilder struct {}

func (b *StdoutChannelBuilder) CreateChannel(conf *ChannelConfig) Channel {
	return &StdoutChannel{conf.Name}		
}

func (sc *MStdoutChannel)Log(m *Message) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	b := m.Formated.String()
	fmt.Print(b)
}

func (sc *MStdoutChannel)Flush() {
}

func (sc *MStdoutChannel)SetProperty(k string, v string) {
	if k == "name" {
		sc.name = v
	}
}
	
func (sc *MStdoutChannel)GetProperty(k string) string {
	if k == "name" {
		return sc.name
	}
	return ""
}

func (sc *MStdoutChannel)GetAllProperty() map[string]string {
	return map[string]string{}
}

func (sc *MStdoutChannel)AddChannel(c Channel) {

}

func (sc *MStdoutChannel)GetAllChannel() []Channel {
	return nil
}

func (sc *MStdoutChannel)GetChannel(name string) Channel {
	return nil
}

type MStdoutChannelBuilder struct {}

func (b *MStdoutChannelBuilder) CreateChannel(conf *ChannelConfig) Channel {
	return &MStdoutChannel{sync.Mutex{}, conf.Name}		
}