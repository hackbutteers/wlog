package wlog


type LogChannel struct {
	ArgsHandler   UserHandler
	LogFormatter  Formatter
	LogLevel      Level
	Name          string
	Sinks         []Channel
}

func NewLogChannel(name string) *LogChannel {
	return &LogChannel{nil, nil, TraceLevel, name, make([]Channel, 0, 5)}
}
func (lc *LogChannel)SetLogChannelName(s string) {
	lc.Name = s
}

func (lc *LogChannel)GetLogChannelName() string {
	return lc.Name
}

func (lc *LogChannel)ShouldLog(l Level) bool {
	return l >= lc.LogLevel
}

func (lc *LogChannel)SetLogLevel(l Level) {
	lc.LogLevel = l
}

func (lc *LogChannel)GetLogLevel() Level {
	return lc.LogLevel
}

func (lc *LogChannel)SetUserHandler(h UserHandler) {
	lc.ArgsHandler = h
}

func (lc *LogChannel)GetUserhandler() UserHandler {
	return lc.ArgsHandler
}

func (lc *LogChannel)SetFormatter(f Formatter) {
	lc.LogFormatter = f
}

func (lc *LogChannel)GetFormatter() Formatter {
	return lc.LogFormatter
}

func (lc *LogChannel)Sink(m *Message) {
	for _, c := range lc.Sinks {
		c.Log(m)
	}	
}
func (lc *LogChannel)LogF(ml *MessageLite, format string, v...interface{}) {
	if !lc.ShouldLog(ml.LogLevel) {
		return
	}
	msg := NewMessage(ml)
	lc.ArgsHandler.Formatf(msg.Raw, format, v...)
	lc.LogFormatter.Format(msg)
	lc.Sink(msg)
}

func (lc *LogChannel)Log(ml *MessageLite, v...interface{}) {
	if !lc.ShouldLog(ml.LogLevel) {
		return
	}
	msg := NewMessage(ml)
	lc.ArgsHandler.Format(msg.Raw, v...)
	lc.LogFormatter.Format(msg)
	lc.Sink(msg)
}

func (lc *LogChannel)AddChannel(c Channel) bool {
	addName := c.GetProperty("name")
	for i := 0; i < len(lc.Sinks); i++ {
		if lc.Sinks[i].GetProperty("name") == addName {
			return false
		}
	}
	lc.Sinks = append(lc.Sinks, c)
	return true
}

func (lc *LogChannel)RemoveChannel(name string) {
	for i := 0; i < len(lc.Sinks); i++ {
		if lc.Sinks[i].GetProperty("name") == name {
			lc.Sinks = append(lc.Sinks[:i],lc.Sinks[:i+1]...)
		}
	}
}