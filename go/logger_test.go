package wlog_test

import (
	"testing"
	"fmt"
	wlog "github.com/hackbutteers/wlog/go"
)
/*
func TestLogger(t *testing.T) {
	lg := wlog.NewLogger("mylog")
	lc := wlog.NewLogChannel("mylogc")
	lc.SetLogLevel(wlog.InfoLevel)
	df := wlog.NewDefaultTestFormatter()
	df.CompileConf("%+")
	lc.SetFormatter(df)
	lc.SetUserHandler(wlog.NewDefaultTextHandler())
	lc.AddChannel(wlog.GetChannel(nil))
	lg.AddChannel("mylogc", lc)
	fmt.Println("start testing logger")	
	lg.Debugf("this should not display")
	lg.Debug("this should not display")

	lg.Info("this should  display")
	lg.Infof("%s", "this should  display")
	fmt.Println("end testing logger")	
}
*/
/*
func TestLogger(t *testing.T) {
	lc, ok := wlog.ConfigLoad("./etc/conf.json")
	if !ok {
		fmt.Println("load error")
	}
	lcc := &lc.Loggers[0]
	lg := wlog.BuildLogger(lcc)
	if lg == nil {
		fmt.Println("log build error")
		return
	}
	fmt.Println("start testing logger")	
	lg.Debugf("this should not display")
	lg.Debug("this should not display")

	lg.Info("this should  display")
	lg.Infof("%s", "this should  display")
	fmt.Println("end testing logger")	
}
*/

func TestLogger(t *testing.T) {
	ok := wlog.LoggerInitFromFile("./etc/conf.json")
	if !ok {
		fmt.Println("init error")
		return
	}

	lg := wlog.GetLogger("mytest_log")
	lg1 := wlog.GetLogger("mytest_log1")
	if lg == nil {
		fmt.Println("log get error")
		return
	}

	if lg1 == nil {
		fmt.Println("log get error")
		return
	}

	fmt.Println("start testing logger")	
	
	lg.Debugf("this should not display")
	lg.Debug("this should not display")

	lg.Info("this should  display")
	lg.Infof("%s", "this should  display")
	
	lg1.Debugf("this should not display")
	lg1.Debug("this should not display")
	for i := 0; i < 100000; i++ {
		lg1.Info("this should  display")
		
		lg1.Infof("%s", "this should  display")
	}

	fmt.Println("end testing logger")	
}