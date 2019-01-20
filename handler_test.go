package glo

import (
	"bytes"
	"testing"
)

func TestHandler(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	expect := formatter.Format(Info, "x") + "\n"
	bfr := bytes.NewBufferString("")

	hndl := NewHandler(bfr).SetFormatter(formatter)

	// Test write
	hndl.Log(Info, "x")
	if rs := bfr.String(); rs != expect {
		t.Errorf("bufio did not receive the log. expected(%#v) got(%#v)", expect, rs)
	}

	// Test SetLevel
	bfr.Truncate(0)
	hndl.SetLevel(Emergency)
	hndl.Log(Debug, "x")
	if rs := bfr.String(); rs != "" {
		t.Error("severity was not ignored")
	}

}
