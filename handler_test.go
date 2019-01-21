package glo

import (
	"bytes"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	formatter := NewFormatter(testFormat)
	expect := formatter.Format(time.Now(), Info, "x") + "\n"
	bfr := bytes.NewBufferString("")

	hndl := NewHandler(bfr).SetFormatter(formatter)

	// Test write
	hndl.Log(Info, "x")
	if rs := bfr.String(); rs != expect {
		t.Errorf("bufio did not receive the log. expected(%#v) got(%#v)", expect, rs)
	}

	// Test SetLevel
	bfr.Truncate(0)
	hndl.PushValidator(NewValidatorLevel(Emergency))
	hndl.Log(Debug, "x")
	hndl.Log(Info, "x")
	hndl.Log(Notice, "x")
	hndl.Log(Warning, "x")
	hndl.Log(Error, "x")
	hndl.Log(Critical, "x")
	hndl.Log(Alert, "x")
	if rs := bfr.String(); rs != "" {
		t.Error("severity was not ignored")
	}
	hndl.Log(Emergency, "x")
	if rs := bfr.String(); rs != "[EMERGENCY] x []\n" {
		t.Error("severity was ignored")
	}

}
