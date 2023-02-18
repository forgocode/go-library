package logger

import (
	"testing"
)

func Test_InitLogger(t *testing.T) {
	opt := Options{
		FileName: "running.log",
		Level:    "debug",
	}
	logger := New(opt)
	logger.Debugln("test debug")
	logger.Infoln("test info")
	logger.Errorln("test error")

	logger.Debugln("test debug")
	logger.Infoln("test info")
	logger.Errorln("test error")

	l := logger.Named("atp")

	l.Debugln("test debug\n")
	l.Infoln("test info\n")
	l.Errorln("test error\n")
	l.Debugf("test debug\n")
	l.Infof("test info\n")
	l.Errorf("test error\n")
	// os.Remove("./running.log")
}
