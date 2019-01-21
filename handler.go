package glo

import (
	"io"
	"time"
)

// Handler logs to stdout
type Handler interface {
	Logger
	SetFormatter(Formatter) Handler
	SetLevel(Level) Handler
}

// NewHandler creates handler that prints to an io.Writer
func NewHandler(writer io.Writer) Handler {
	return &handler{
		writer,
		NewFormatter(DefaultFormat),
		Debug,
	}
}

type handler struct {
	writer    io.Writer
	formatter Formatter
	level     Level
}

// Log logs a line with a specific level
func (h *handler) Log(level Level, line string, params ...interface{}) error {
	if level < h.level {
		return nil
	}
	l := h.formatter.Format(time.Now(), level, line, params...) + "\n"
	_, err := h.writer.Write([]byte(l))
	return err
}

func (h *handler) SetFormatter(formatter Formatter) Handler {
	h.formatter = formatter
	return h
}

func (h *handler) SetLevel(level Level) Handler {
	h.level = level
	return h
}
