package glo

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatter(t *testing.T) {
	formatter := NewFormatter(DefaultFormat)
	now := time.Now()
	timestamp := now.Format(dateFormat)
	line := formatter.Format(now, Info, "%-10s %-10d", "pad me", 123)
	if line != fmt.Sprintf("%s [INFO] pad me     123       ", timestamp) {
		t.Error("failed to format correctly")
	}
}
