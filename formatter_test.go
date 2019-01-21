package glo

import (
	"testing"
	"time"
)

const testFormat = "[%[2]s] %[3]s %[4]#"

func TestFormatter(t *testing.T) {
	formatter := NewFormatter(testFormat)
	line := formatter.Format(time.Now(), Info, "x")
	if line != "[INFO] x []" {
		t.Error("failed to format correctly")
	}
}
