package wlog

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)
type ChannelConfig struct {
	Ctype            string
	Name             string
	Dir              string
	Fix              string
	Flush_size       int
	Max_files        int
	Singal_file_size int
	Total_size       int
	Save_days        int
	Flush_sec        int
	Raw              string 
}

type LogChannelConfig struct {
	Name        string
	Level       string
	Handlername string
	Formatter   string
	Format      string
	Time_utc    string
	Sinks       []ChannelConfig
}
type LogConfig struct {
	Name string
	Channels []LogChannelConfig
}

type LogALlConfig struct {
	Loggers []LogConfig
}
func ConfigLoad(path string) (*LogALlConfig, bool) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, false
	}
	var lc LogALlConfig
	e := json.Unmarshal(data, &lc)
	if e != nil {
		fmt.Println(e.Error())
		fmt.Println(" parse json error")
		return nil, false
	}
	return &lc, true
}
