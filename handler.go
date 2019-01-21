package glo

import (
	"io"
	"time"
)

// Handler logs to stdout
type Handler interface {
	Logger
	SetFormatter(Formatter) Handler
	ClearValidators() Handler
	PushValidator(Validator) Handler
}

// NewHandler creates handler that prints to an io.Writer
func NewHandler(writer io.Writer) Handler {
	return &handler{
		writer,
		NewFormatter(DefaultFormat),
		[]Validator{},
	}
}

type handler struct {
	writer     io.Writer
	formatter  Formatter
	validators []Validator
}

// Log logs a line with a specific level
func (h *handler) Log(level Level, line string, params ...interface{}) error {
	valid := true
	for _, v := range h.validators {
		if !v.Validate(level, line, params) {
			valid = false
			break
		}
	}
	if !valid {
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

func (h *handler) ClearValidators() Handler {
	h.validators = []Validator{}
	return h
}

func (h *handler) PushValidator(validator Validator) Handler {
	h.validators = append(h.validators, validator)
	return h
}
