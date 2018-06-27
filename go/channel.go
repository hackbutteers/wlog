package wlog 

type Channel interface {
	
	Log(m *Message)

	Flush()

	SetProperty(k string, v string)
	
	GetProperty(k string) string
	
	GetAllProperty() map[string]string
	
	AddChannel(c Channel)
	
	GetAllChannel() []Channel
	
	GetChannel(name string) Channel
}
