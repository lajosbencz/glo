package glo

// Logger logs a line with a specific level
type Logger interface {
	Handler
	LoggerDebug
	LoggerInfo
	LoggerNotice
	LoggerWarning
	LoggerError
	LoggerCritical
	LoggerAlert
	LoggerEmergency
	ClearHandlers() Logger
	AddHandler(Handler) Logger
}

// NewLogger instantiates a Logger
func NewLogger() Logger {
	l := &logger{
		handlers: []Handler{},
	}
	return l
}

type logger struct {
	handlers []Handler
}

// Log sends output to all registered handlers
func (l logger) Log(level Level, line string, params ...interface{}) error {
	var err1 error
	for _, hndl := range l.handlers {
		if err := hndl.Log(level, line, params...); err != nil && err1 == nil {
			err1 = err
		}
	}
	return err1
}

func (l *logger) ClearHandlers() Logger {
	// @TODO loop and close
	l.handlers = []Handler{}
	return l
}

func (l *logger) AddHandler(h Handler) Logger {
	l.handlers = append(l.handlers, h)
	return l
}
