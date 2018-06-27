package wlog_test
/*
import (
	"os"
	"fmt"
	wlog "github.com/hackbutteers/wlog/go"
	"testing"
	"bytes"
)

func TestUserHandler(t *testing.T) {
	dth := wlog.NewDefaultTextHandler()
	//dth :=&wlog.DefaultTextHandler{"11111"}
	var b bytes.Buffer
	dth.Formatf(&b, "%s-%d\n", "fmtf", 1)
	fmt.Fprint(os.Stdout,&b)
	var bb bytes.Buffer
	dth.Format(&bb, "fmtf-", 1, "\n")
	fmt.Fprint(os.Stdout,&bb)
}*/