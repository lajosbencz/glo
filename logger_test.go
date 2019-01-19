package glo

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	expect := formatter.Format(Info, "test") + "\n"
	bfr := bytes.NewBufferString("")
	bfr2 := bytes.NewBufferString("")

	log := NewLogger()

	// Test write
	log.AddHandler(NewHandler(bfr, formatter))
	log.AddHandler(NewHandler(bfr2, formatter))
	log.Info("test")
	if rs := bfr.String(); rs != expect {
		t.Errorf("bfr expected(%#v) got(%#v)", expect, rs)
	}
	if rs := bfr2.String(); rs != expect {
		t.Errorf("bfr2 expected(%#v) got(%#v)", expect, rs)
	}

	// Test severity
	bfr.Truncate(0)
	log.Debug("x")
	log.Info("x")
	log.Notice("x")
	log.Warning("x")
	log.Error("x")
	log.Critical("x")
	log.Alert("x")
	log.Emergency("x")
	if rs := bfr.String(); rs != "[DEBUG] x []\n[INFO] x []\n[NOTICE] x []\n[WARNING] x []\n[ERROR] x []\n[CRITICAL] x []\n[ALERT] x []\n[EMERGENCY] x []\n" {
		t.Errorf("bufio has wrong value. got(%#v)", rs)
	}

	// Test clear
	bfr.Truncate(0)
	log.ClearHandlers()
	log.Info("test")
	if rs := bfr.String(); rs != "" {
		t.Errorf("bufio was not empty. got(%#v)", rs)
	}

	// Test first handler error
	log.AddHandler(&mockHandler{
		handler{bfr, formatter},
	})
	if err := log.Log(Debug, "err"); err != mockErrHandler {
		t.Error("expecting error")
	}
}

type mockHandler struct {
	handler
}

func (h mockHandler) Log(level Level, line string, params ...interface{}) error {
	return mockErrHandler
}

type mockHandlerError string

func (h mockHandlerError) Error() string {
	return string(h)
}

var mockErrHandler mockHandlerError = "err"
