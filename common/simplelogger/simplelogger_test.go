package simplelogger

import "testing"

func Test_ttt(t *testing.T) {
	log := GetLogger()
	log.Info("info")
	log.Debug("debug")
	log.Error("error")
}
