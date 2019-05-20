package simplelogger

/*
Usage:
**********************************************************************
var (
    log = logger.GetLogger()
)

func init() {
    logfile := "C:/Users/cc/Desktop/small-site-main.log"
    output, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        log.Error("Failed to log to file, using default os.Stdout")
        output = os.Stdout
    }
    log.SetOutput(output)
}
**********************************************************************
*/

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"runtime"
)

type CcLogger struct {
	Logger *logrus.Logger
}

func getCallerInfo() string {
	pc, filename, _, _ := runtime.Caller(2)
	return "[" + filename + "] => [" + runtime.FuncForPC(pc).Name() + "]:"
}

func (l CcLogger) Info(s string) {
	l.Logger.Info(s)
}

func (l CcLogger) Debug(s string) {
	fmt.Println(getCallerInfo())
	l.Logger.Debug(s)
}

func (l CcLogger) Error(s string) {
	fmt.Println(getCallerInfo())
	l.Logger.Error(s)
}

func (l CcLogger) Panic(s string) {
	fmt.Println(getCallerInfo())
	panic(s)
}

var (
	_layoutDatetime = "2006-01-02 15:04:05"
	logger          = CcLogger{}
)

func init() {
	logger.Logger = &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: _layoutDatetime,
			//defaultLogFormat = "[%lvl%]: %time% - %msg%"
			LogFormat: "%time% - [%lvl%] - %msg%\n",
		},
	}
}

func GetLogger() CcLogger {
	return logger
}
