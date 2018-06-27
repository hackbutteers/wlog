package wlog

import (
	"fmt"
)

type FlagFormatter interface {
	format(m *Message)
}


type FlagFormatter_v struct {}

func (v *FlagFormatter_v)format(m *Message) {
	m.Raw.WriteTo(m.Formated)
}

type FlagFormatter_t struct {}

func (v *FlagFormatter_t)format(m *Message) {
	//go do not have thread id	
}

type FlagFormatter_n struct{}

func (v *FlagFormatter_n)format(m *Message) {
	m.Formated.WriteString(*m.LogName)

}

type FlagFormatter_l struct{}

func (v *FlagFormatter_l)format(m *Message) {
	m.Formated.WriteString(LevelToString(m.LogLevel))
}

type FlagFormatter_L struct{}

func (v *FlagFormatter_L)format(m *Message) {
	m.Formated.WriteString(LevelToStringShort(m.LogLevel))
}

type FlagFormatter_T struct{}

func (v *FlagFormatter_T)format(m *Message) {
	m.Formated.WriteString(formatTime(m.Time))
}

type FlagFormatter_char struct {
	c rune
}

func (v *FlagFormatter_char)format(m *Message) {
	m.Formated.WriteRune(v.c)
}

type FlagFormatter_str struct {
	str string	
}

func NewFlagFormatter_str(str string) *FlagFormatter_str {
	return &FlagFormatter_str{str}
}
func (v *FlagFormatter_str)format(m *Message) {
	m.Formated.WriteString(v.str)
}

type FlagFormatter_default struct {}

func (v *FlagFormatter_default)format(m *Message) {
	fmt.Fprintf(m.Formated, "%s %s %s ", formatTime(m.Time), *m.LogName, LevelToString(m.LogLevel))
	m.Raw.WriteTo(m.Formated)
}

func FlagFormat(f FlagFormatter, m *Message) {
	f.format(m)
}