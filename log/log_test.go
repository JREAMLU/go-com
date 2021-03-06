package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	SetLevel(Ldebug)
	SetLogDir(".")
	Printf("Print: foo\n")
	Print("Print: foo")

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")

	SetLevel(Lerror)

	Printf("Print: foo\n")
	Print("Print: foo")

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")
	WaitFlush()
}
