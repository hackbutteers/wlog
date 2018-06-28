package wlog

import (
	"bytes"
)

type DefaultTestFormatter struct {
	Utc bool
	ffs []FlagFormatter
}

func NewDefaultTestFormatter() *DefaultTestFormatter {
	return &DefaultTestFormatter{false, make([]FlagFormatter, 0, 10)}
}
func (f *DefaultTestFormatter)handleString(str string){
	if len(str) > 0 {
		s := NewFlagFormatter_str(str)
		f.ffs = append(f.ffs, s)
	}
}

func (f *DefaultTestFormatter)handleChar(c rune) bool {
	switch c {
	case '+':
		flag := &FlagFormatter_default{}
		f.ffs = append(f.ffs, flag)
		return true
	case 'v':
		flag := &FlagFormatter_v{}
		f.ffs = append(f.ffs, flag)
		return true
	case 'l':
		flag := &FlagFormatter_l{}
		f.ffs = append(f.ffs, flag)
		return true
	case 'L':
		flag := &FlagFormatter_L{}
		f.ffs = append(f.ffs, flag)
		return true
	case 'T':
		flag := &FlagFormatter_T{}
		f.ffs = append(f.ffs, flag)
		return true
	case 't':
		flag := &FlagFormatter_t{}
		f.ffs = append(f.ffs, flag)
		return true
	case 'n':
		flag := &FlagFormatter_n{}
		f.ffs = append(f.ffs, flag)
		return true
	}
	return false
}

func (f *DefaultTestFormatter)CompileConf(str string) bool {
	var buffer bytes.Buffer
	buffer.Reset()
	isflag := false
	for _, c := range str {
		if c == '%' {
			if isflag {
				buffer.WriteRune(c)
				continue
			}
			isflag = true
			continue
		}

		if isflag {
			f.handleString(buffer.String())
			buffer.Reset()
			r := f.handleChar(c)
			isflag = false
			if r == false  {
				return false
			}
		} else {
			buffer.WriteRune(c)
		}

	}
	f.handleString(buffer.String())
	return true;
}

func (f *DefaultTestFormatter)Format(m *Message) {
	for _, f := range f.ffs {
		FlagFormat(f, m)
	}
	m.Formated.WriteRune('\n')
}

type DefaultTestFormatterBuilder struct {}

func (b *DefaultTestFormatterBuilder) CreateFormatter(conf *LogChannelConfig) Formatter {
	return NewDefaultTestFormatter()		
}