package glo

import (
	"bytes"
	"testing"
)

func TestHandler(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	expect := formatter.Format(Info, "test")
	bfr := bytes.NewBufferString("")

	hndl := NewHandler(bfr, formatter)
	hndl.Log(Info, "test")

	if rs, _ := bfr.ReadString('\n'); rs != expect {
		t.Errorf("bufio did not receive the log. expected(%#v) got(%#v)", expect, rs)
	}
}
