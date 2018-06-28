package wlog

import (
	"fmt"
	"time"
	"strconv"
)
type MaxSizeFileChannel struct {
	max_size  int
	dir       string
	prefix    string
	fix       string
	name      string
	fileName  string
	seq       int
	lf        *LogFile
}

func (sc *MaxSizeFileChannel)genFileName() string {
	ret := sc.dir + "/" 
	ret = ret + sc.prefix
	ret = ret + formatTimeForFile(time.Now())
	ret = ret + strconv.Itoa(sc.seq)
	ret = ret + sc.fix
	sc.seq++
	return ret
}
func (sc *MaxSizeFileChannel)rotate() {
	sc.fileName = sc.genFileName()
	sc.lf.Reset(sc.fileName)
}
func (sc *MaxSizeFileChannel)Log(m *Message) {
	n := sc.lf.WriteSize()
	if m.Formated.Len() + n > sc.max_size {
		sc.rotate()
	}
	sc.lf.WriteBuffer(m.Formated)
}

func (sc *MaxSizeFileChannel)Flush() {
}

func (sc *MaxSizeFileChannel)SetProperty(k string, v string) {
	if k == "name" {
		sc.name = v
	}
}
	
func (sc *MaxSizeFileChannel)GetProperty(k string) string {
	if k == "name" {
		return sc.name
	}
	return ""
}

func (sc *MaxSizeFileChannel)GetAllProperty() map[string]string {
	return map[string]string{}
}

func (sc *MaxSizeFileChannel)AddChannel(c Channel) {

}

func (sc *MaxSizeFileChannel)GetAllChannel() []Channel {
	return nil
}

func (sc *MaxSizeFileChannel)GetChannel(name string) Channel {
	return nil
}

type MaxSizeFileChannelBuilder struct {}

func (b *MaxSizeFileChannelBuilder) CreateChannel(conf *ChannelConfig) Channel {
	r :=  &MaxSizeFileChannel{}	
	r.dir = conf.Dir
	r.fix = conf.Fix
	r.name = conf.Name
	r.max_size = conf.Total_size
	r.seq = 0
	r.prefix = conf.Prefix	
	r.fileName = r.genFileName()
	lf, e := NewLogFile(r.fileName)
	if e != nil {
		fmt.Println("dir not exit")
		return nil
	}
	r.lf = lf
	return r
}
