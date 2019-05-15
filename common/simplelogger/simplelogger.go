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
	"github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

var (
	_layoutDatetime = "2006-01-02 15:04:05"
	log             = logrus.New()
)

func init() {
	log = &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: _layoutDatetime,
			//defaultLogFormat = "[%lvl%]: %time% - %msg%"
			LogFormat: "%time% - [%lvl%] - %msg%\n",
		},
	}
}

func GetLogger() *logrus.Logger {
	return log
}
