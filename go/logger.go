package wlog 

import (
	"time"
)


type Logger struct {
	name  string
	sinks map[string]*LogChannel
}

func NewLogger(name string) *Logger {
	return &Logger{name, make(map[string]*LogChannel, 5)}
}

func (lg *Logger) Name() string {
	return lg.name
}

// do not lock this it will call in a init thread
func (lg *Logger) AddChannel(name string, lc *LogChannel) {
	_, ok := lg.sinks[name]
	if ok {
		return
	}
	lg.sinks[name] = lc
}
//log operations
func (l *Logger) log(lvl Level, v...interface{}) {
	ml := NewMessageLite(&l.name, time.Now(), lvl)
	for _, lc := range l.sinks {
		lc.Log(ml, v...)
	}
}

func (l *Logger) logf(lvl Level, f string, v...interface{}) {
	ml := NewMessageLite(&l.name, time.Now(), lvl)
	for _, lc := range l.sinks {
		lc.LogF(ml, f, v...)
	}
}

func (l *Logger) Trace(v...interface{}) {
	l.log(TraceLevel, v...)
}

func (l *Logger) Tracef(f string, v...interface{}) {
	l.logf(TraceLevel, f, v...)
}

func (l *Logger) Debug(v...interface{}) {
	l.log(DebugLevel, v...)
}

func (l *Logger) Debugf(f string, v...interface{}) {
	l.logf(DebugLevel, f, v...)
}

func (l *Logger) Info(v...interface{}) {
	l.log(InfoLevel, v...)
}

func (l *Logger) Infof(f string, v...interface{}) {
	l.logf(InfoLevel, f, v...)
}

func (l *Logger) Warn(v...interface{}) {
	l.log(WarnLevel, v...)
}

func (l *Logger) Warnf(f string, v...interface{}) {
	l.logf(WarnLevel, f, v...)
}

func (l *Logger) Error(v...interface{}) {
	l.log(ErrorLevel, v...)
}

func (l *Logger) Errorf(f string, v...interface{}) {
	l.logf(ErrorLevel, f, v...)
}

func (l *Logger) Fatal(v...interface{}) {
	l.log(FatalLevel, v...)
}

func (l *Logger) Fatalf(f string, v...interface{}) {
	l.logf(FatalLevel, f, v...)
}