package glo

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"testing"
	"time"
)

func TestFacility(t *testing.T) {
	formatter := NewFormatter("[{L}] {M}")
	expect := formatter.Format(time.Now(), Info, "x") + "\n"
	bfr := bytes.NewBufferString("")
	bfr2 := bytes.NewBufferString("")

	log := NewFacility()

	// Test write
	log.PushHandler(NewHandler(bfr).SetFormatter(formatter))
	log.PushHandler(NewHandler(bfr2).SetFormatter(formatter))
	log.Log(Info, "x")
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
	if rs := bfr.String(); rs != "[DEBUG] x\n[INFO] x\n[NOTICE] x\n[WARNING] x\n[ERROR] x\n[CRITICAL] x\n[ALERT] x\n[EMERGENCY] x\n" {
		t.Errorf("bufio has wrong value. got(%#v)", rs)
	}

	// Test clear
	bfr.Truncate(0)
	log.ClearHandlers()
	log.Info("x")
	if rs := bfr.String(); rs != "" {
		t.Errorf("bufio was not empty. got(%#v)", rs)
	}

	// Test first handler error
	log.PushHandler(&mockHandler{
		handler{bfr, formatter, []Filter{}},
	})
	if err := log.Log(Debug, "x"); err != mockErrHandler {
		t.Error("expecting error")
	}
}

func TestStdFacility(t *testing.T) {
	o, e := captureStdout(useStdFacility)

	rgxo := regexp.MustCompile(`\[INFO\] x\n$`)
	rgxe := regexp.MustCompile(`\[ERROR\] x\n$`)
	if !rgxo.MatchString(o) {
		t.Error("info was not written to stdout")
	}
	if !rgxe.MatchString(e) {
		t.Error("error was not written to stderr")
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

func captureStdout(f func()) (string, string) {
	oldo := os.Stdout
	ro, wo, _ := os.Pipe()
	os.Stdout = wo
	olde := os.Stderr
	re, we, _ := os.Pipe()
	os.Stderr = we

	f()

	wo.Close()
	os.Stdout = oldo
	var bufo bytes.Buffer
	io.Copy(&bufo, ro)

	we.Close()
	os.Stderr = olde
	var bufe bytes.Buffer
	io.Copy(&bufe, re)

	return bufo.String(), bufe.String()
}

func useStdFacility() {
	log := NewStdFacility()
	log.Info("x")
	log.Error("x")
}
