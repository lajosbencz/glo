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

	// Simple write
	hndl.Log(Info, "x")
	if rs := bfr.String(); rs != expect {
		t.Error("log was not written")
	}

	// Validate Level
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

	// Clear validators
	bfr.Truncate(0)
	hndl.ClearValidators()
	hndl.Log(Debug, "x")
	if rs := bfr.String(); rs != "[DEBUG] x []\n" {
		t.Error("validators were not cleared")
	}
}
