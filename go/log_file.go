package wlog 

import (
	"bytes"
	"os"
	"time"
)

type LogFile struct {
	file      *os.File
	name       string
	create     time.Time 
	write      int 
}

func NewLogFile(name string) (*LogFile, error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}
	return &LogFile{f, name, time.Now(), 0}, nil
}

func (f *LogFile) WriteSize() int {
	return f.write
}

func (f *LogFile) CreateTime() time.Time  {
	return f.create
}
func (f *LogFile)Reset(name string) (*LogFile, error) {
	f.file.Close()
	fp, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	f.file = fp
	f.name = name
	f.create = time.Now()
	f.write = 0
	return f, nil
} 

func (f *LogFile) WriteBuffer(buffer *bytes.Buffer) int {
	len := buffer.Len()
	b := buffer.Bytes()
	wlen, e := f.file.Write(b)
	if e != nil {
		return 0
	}
	if len == wlen {
		f.write += len
		return len
	}
	return 0;
}

func (f *LogFile)Flush() {
	f.file.Sync()
}

func (f *LogFile)Close() {
	f.file.Close()
}
