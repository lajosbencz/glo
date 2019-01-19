package glo

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	expect := formatter.Format(Info, "test") + "\n"
	bfr := bytes.NewBufferString("")

	log := NewLogger()
	log.AddHandler(NewHandler(bfr, formatter))

	log.Info("test")

	if rs, _ := bfr.ReadString('\n'); rs != expect {
		t.Errorf("bufio did not receive the log. expected(%#v) got(%#v)", expect, rs)
	}
}
