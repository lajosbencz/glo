package glo

import (
	"bytes"
	"testing"
)

func TestHandler(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	expect := formatter.Format(Info, "test") + "\n"
	bfr := bytes.NewBufferString("")

	hndl := NewHandler(bfr, formatter)
	hndl.Log(Info, "test")

	if rs := bfr.String(); rs != expect {
		t.Errorf("bufio did not receive the log. expected(%#v) got(%#v)", expect, rs)
	}
}
