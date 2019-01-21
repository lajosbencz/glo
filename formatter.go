package glo

import (
	"fmt"
	"time"
)

const (
	// DefaultFormat is used by default
	DefaultFormat string = "%[1]s [%[2]s] %[3]s %[4]v"
)

// Formatter formats a log event
type Formatter interface {
	Format(time.Time, Level, string, ...interface{}) string
}

// NewFormatter creates a Formatter from a string
func NewFormatter(f string) Formatter {
	return &formatter{f}
}

type formatter struct {
	format string
}

func (f *formatter) Format(time time.Time, level Level, line string, params ...interface{}) string {
	return fmt.Sprintf(f.format, time.Format("2006-01-02T15:04:05Z07:00"), level, line, params)
}
