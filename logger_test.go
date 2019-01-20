package glo

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := &mockLogger{
		Debug,
		"",
		[]interface{}{},
	}
	logger.Log(Alert, "x", 1, 2)
	if logger.mLevel != Alert {
		t.Error("level")
	}
	if logger.mLine != "x" {
		t.Error("line")
	}
	// todo
	// if logger.mParams[1].(int) != 1 || logger.mParams[2].(int) != 2 {
	// 	t.Error("params")
	// }
}

type mockLogger struct {
	mLevel  Level
	mLine   string
	mParams []interface{}
}

func (l *mockLogger) Log(level Level, line string, params ...interface{}) error {
	l.mLevel = level
	l.mLine = line
	l.mParams = params
	return nil
}
