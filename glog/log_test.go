package glog

import (
	"testing"
	"time"
)

func TestLocal(t *testing.T) {
	tm := time.Now()
	t.Log(tm.Zone())
}

func TestLevel(t *testing.T) {
	InitLogger(DEV, nil)
	SetLevel(WarnLevel)
	if Level() != WarnLevel {
		t.Fail()
		t.Log(_logger.Level())
		c := _logger.(*console)
		t.Log(c.level)
	}
}

func testConsoleLog(t *testing.T) {
	InitLogger(DEV, nil)
	Debug("this is a debug info\n")
	Info("this is a info %s", "logger init successfully.\n")
	Warn("this is a warning: base value should not be %d\n", 0)
	Error("this is a error log\n")

	SetPrefix(InfoLevel, "information:")
	Info("modify info log prefix to %s\n", Prefix(InfoLevel))
	Info("no new line")
	Info("new prefix info log\n")
}

func testFileLog(t *testing.T) {
	InitLogger(PRO, map[string]interface{}{"typ": "file", "dir": "./logs"})
	//InitLogger(PRO, map[string]interface{}{"typ": "file", "seconds": 15})
	Debug("this is a debug info\n")
	Info("this is a info %s", "logger init successfully.\n")
	Warn("this is a warning: base value should not be %d\n", 0)
	Error("this is a error log\n")

	SetPrefix(InfoLevel, "information:")
	Info("modify info log prefix to %s\n", Prefix(InfoLevel))
	Info("no new line")
	Info("new prefix info log\n")

	time.Sleep(3 * time.Second)
	Info("log after 3 seconds.")

	time.Sleep(5 * time.Second)
	Info("log after 5 seconds.")

	time.Sleep(8 * time.Second)
	Info("log after 8 seconds.")

	time.Sleep(6 * time.Second)
	Info("log after 6 seconds.")

	time.Sleep(6 * time.Second)
	Info("log after 6 seconds.")

	Close()
}

func testFileLogClean(t *testing.T) {
	InitLogger(PRO, map[string]interface{}{"typ": "file", "dir": "./logs", "suffix": "-{{yyyy}}{{mm}}{{dd}}"})
	//InitLogger(PRO, map[string]interface{}{"typ": "file", "seconds": 5})
	Debug("this is a debug info\n")
	Info("this is a info %s", "logger init successfully.\n")
	Warn("this is a warning: base value should not be %d\n", 0)
	Error("this is a error log\n")

	SetPrefix(InfoLevel, "information:")
	Info("modify info log prefix to %s\n", Prefix(InfoLevel))
	Info("no new line")
	Info("new prefix info log\n")

	time.Sleep(3 * time.Second)
	Info("log after 3 seconds.")

	time.Sleep(5 * time.Second)
	Info("log after 5 seconds.")

	time.Sleep(8 * time.Second)
	Info("log after 8 seconds.")

	time.Sleep(6 * time.Second)
	Info("log after 6 seconds.")

	time.Sleep(60 * time.Second)
	Info("log after 6 seconds.")
	Close()
}

/*
func TestNsqLog(t *testing.T) {
	InitLogger(PRO, map[string]interface{}{"typ": "nsq", "nsqdAddr": "10.10.133.80:4150"})

	Debug("this is a debug info\n")
	Info("this is a info %s", "logger init successfully.\n")
	Warn("this is a warning: base value should not be %d\n", 0)
	Error("this is a error log\n")

	SetPrefix(InfoLevel, "information:")
	Info("modify info log prefix to %s\n", Prefix(InfoLevel))
	Info("no new line")
	Info("new prefix info log\n")

	time.Sleep(3 * time.Second)
	Info("log after 3 seconds.")

	time.Sleep(5 * time.Second)
	Info("log after 5 seconds.")

	time.Sleep(8 * time.Second)
	Info("log after 8 seconds.")

	time.Sleep(6 * time.Second)
	Info("log after 6 seconds.")

	Close()
}
*/
