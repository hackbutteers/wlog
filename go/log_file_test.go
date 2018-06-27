package wlog_test
/*
import (
	"bytes"
	"fmt"
	"testing"
	wlog "github.com/hackbutteers/wlog/go"
)

func TestLogFile(t *testing.T) {
	lf, err := wlog.NewLogFile("./testing.log");
	if	err != nil {
		fmt.Println("error")
	}
	var b bytes.Buffer
	s := "my log file test"
	b.WriteString(s);
	lf.WriteBuffer(&b)
	fmt.Println("string size", len(s))
	fmt.Println("log file size", lf.WriteSize())
	lf.Reset("./testing.log")
	fmt.Println(lf.WriteSize())
	lf.Close()
}
*/