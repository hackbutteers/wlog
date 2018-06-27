package wlog 

import (
	"fmt"
	"sync"
)


func BuildLogger(lc *LogConfig) *Logger {
	if lc == nil {
		return nil
	}
	lg := NewLogger(lc.Name)
	for _, lcc := range lc.Channels {
		lc := NewLogChannel(lcc.Name)
		l,e := ParseLevel(lcc.Level)
		if e != nil {
			fmt.Println("log level parse error")
			return nil
		}
		lc.SetLogLevel(l)
		f := GetFormatter(&lcc)
		if f == nil {
			fmt.Println("get formatter error")
			return nil
		}
		ok := f.CompileConf(lcc.Format)
		if !ok {
			fmt.Println("compile log parttern eror")
			return nil	
		}	
		lc.SetFormatter(f)
		uh := GetUH(&lcc)
		if uh == nil {
			fmt.Println("get user handler error")
			return nil
		}
		lc.SetUserHandler(uh)
		i := 0
		for _, sc := range lcc.Sinks {
			c := GetChannel(&sc)
			if c == nil {
				fmt.Println("get channel error")
				return nil
			}
			ok := lc.AddChannel(c)
			if !ok {
				fmt.Println("add channel error")
				return nil	
			}
			fmt.Println(lc.Name, "add channel: ", i)
			i++
		}
		lg.AddChannel(lcc.Name, lc)
	}
	return lg
}

type LogRegister struct {
	mu      sync.Mutex
	loggers map[string] *Logger
}

var g_log_register *LogRegister

var g_log_register_once sync.Once

func getRegister() *LogRegister {
	g_log_register_once.Do(func() {
        g_log_register = &LogRegister {sync.Mutex{}, make(map[string] *Logger, 10)}
    })
    return g_log_register
}

func LoggerInitFromFile(path string) bool {
	lac, ok := ConfigLoad(path)
	if !ok {
		fmt.Println("load error")
	}
	getRegister().mu.Lock()
	defer getRegister().mu.Unlock()
	for _, lc := range lac.Loggers {
		_, ok := g_log_register.loggers[lc.Name]
		if ok {
			continue;
		}
		lg := BuildLogger(&lc)
		if lg == nil {
			fmt.Println("create log error")
			return false
		}
		getRegister().loggers[lc.Name] = lg
	}
	return true
}

func GetLogger(name string) *Logger{
	getRegister().mu.Lock()
	defer getRegister().mu.Unlock()
	lg, _ := getRegister().loggers[name]
	return lg
}