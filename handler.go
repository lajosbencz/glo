package glo

import (
	"io"
)

// Handler logs to stdout
type Handler interface {
	Log(level Level, line string, params ...interface{}) error
}

// NewHandler creates handler that prints to an io.Writer
func NewHandler(w io.Writer, f Formatter) Handler {
	return &handler{
		writer:    w,
		formatter: f,
	}
}

type handler struct {
	writer    io.Writer
	formatter Formatter
}

// Log logs a line with a specific level
func (h handler) Log(level Level, line string, params ...interface{}) error {
	l := h.formatter.Format(level, line, params...) + "\n"
	_, err := h.writer.Write([]byte(l))
	return err
}
