package glo

import (
	"testing"
)

func TestFormatter(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	line := formatter.Format(Info, "test")
	if line != "[INFO] test []" {
		t.Error("failed to format correctly")
	}
}
