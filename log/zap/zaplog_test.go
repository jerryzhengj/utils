package zap

import ("testing")


func TestSetLevel(t *testing.T) {
	SetLevel("info")
}

func TestInfo(t *testing.T) {
	Info("test func TestInfo")
}

func TestDebugf(t *testing.T) {
	SetLevel("info")
	Debugf("log level is %s","info")
	SetLevel("debug")
	Debugf("log level is %s","debug")
}
