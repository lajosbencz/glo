package glo

import "fmt"

const (
	// DefaultFormat is used by default
	DefaultFormat string = "[%[1]s] %[2]s %[3]v"
)

// Formatter formats a log event
type Formatter interface {
	Format(Level, string, ...interface{}) string
}

// NewFormatter creates a Formatter from a string
func NewFormatter(f string) Formatter {
	return &formatter{f}
}

type formatter struct {
	format string
}

func (f *formatter) Format(level Level, line string, params ...interface{}) string {
	return fmt.Sprintf(f.format, level, line, params)
}
