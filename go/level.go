package wlog

import (
	"fmt"
	"strings"
)

type Level uint32


const (
	TraceLevel Level = iota
	DebugLevel 
	InfoLevel  
	WarnLevel  
	ErrorLevel 
	FatalLevel 
)

var levels = []Level {
	TraceLevel,
	DebugLevel,
	InfoLevel,
	WarnLevel,
	ErrorLevel,
	FatalLevel,
}

func LevelToString(l Level) string {
	switch l {
	case TraceLevel:
		return "trace";
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}
	return "unkown"
}

func LevelToStringShort(l Level) string {
	switch l {
	case TraceLevel:
		return "T";
	case DebugLevel:
		return "D"
	case InfoLevel:
		return "I"
	case WarnLevel:
		return "W"
	case ErrorLevel:
		return "E"
	case FatalLevel:
		return "F"
	}
	return "U"
}

func LevelToStringUpper(l Level) string {
	switch l {
	case TraceLevel:
		return "TRACE";
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	}
	return "UNKOWN"
}

func ParseLevel(s string) (Level, error) {
	if len(s) > 1 {
		switch strings.ToLower(s) {
		case "trace":
			return TraceLevel, nil
		case "debug":
			return DebugLevel, nil
		case "info":
			return InfoLevel, nil
		case "warn":
			return WarnLevel, nil
		case "error":
			return ErrorLevel, nil
		case "fatal":
			return FatalLevel, nil
		}
	} else if len(s) == 1 {
		switch strings.ToLower(s) {
		case "T":
			return TraceLevel, nil
		case "D":
			return DebugLevel, nil
		case "I":
			return InfoLevel, nil
		case "W":
			return WarnLevel, nil
		case "E":
			return ErrorLevel, nil
		case "F":
			return FatalLevel, nil
		}
	}
	var l Level
	return l, fmt.Errorf("not a valid wlog Level: %q", s)
}