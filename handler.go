package glo

import (
	"io"
)

// Handler logs to stdout
type Handler interface {
	Logger
	SetFormatter(Formatter) Handler
}

// NewHandler creates handler that prints to an io.Writer
func NewHandler(writer io.Writer) Handler {
	return &handler{
		writer,
		NewFormatter(DefaultFormat),
	}
}

type handler struct {
	writer    io.Writer
	formatter Formatter
}

// Log logs a line with a specific level
func (h *handler) Log(level Level, line string, params ...interface{}) error {
	l := h.formatter.Format(level, line, params...) + "\n"
	_, err := h.writer.Write([]byte(l))
	return err
}

func (h *handler) SetFormatter(formatter Formatter) Handler {
	h.formatter = formatter
	return h
}
