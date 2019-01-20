package glo

import (
	"testing"
)

func TestFormatter(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	line := formatter.Format(Info, "x")
	if line != "[INFO] x []" {
		t.Error("failed to format correctly")
	}
}
