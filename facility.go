package glo

// Facility is the main entry when used
type Facility interface {
	Logger
	LoggerDebug
	LoggerInfo
	LoggerNotice
	LoggerWarning
	LoggerError
	LoggerCritical
	LoggerAlert
	LoggerEmergency
	ClearHandlers() Facility
	PushHandler(Handler) Facility
}

// NewFacility creates a logger facility, this is the main entry
func NewFacility() Facility {
	return &facility{
		[]Handler{},
	}
}

type facility struct {
	handlers []Handler
}

// ClearHandlers removes all registered handlers
func (f *facility) ClearHandlers() Facility {
	f.handlers = []Handler{}
	return f
}

// PushHandlers adds a new handler
func (f *facility) PushHandler(h Handler) Facility {
	f.handlers = append(f.handlers, h)
	return f
}

// Log sends output to all registered handlers
func (f *facility) Log(level Level, line string, params ...interface{}) error {
	var err1 error
	for _, hndl := range f.handlers {
		err := hndl.Log(level, line, params...)
		if err != nil && err1 == nil {
			err1 = err
		}
	}
	return err1
}

// Debug logs a debug line
func (f *facility) Debug(line string, params ...interface{}) error {
	return f.Log(Debug, line, params...)
}

// Info logs an info line
func (f *facility) Info(line string, params ...interface{}) error {
	return f.Log(Info, line, params...)
}

// Notice logs a notice line
func (f *facility) Notice(line string, params ...interface{}) error {
	return f.Log(Notice, line, params...)
}

// Warning logs a warning line
func (f *facility) Warning(line string, params ...interface{}) error {
	return f.Log(Warning, line, params...)
}

// Error logs an error line
func (f *facility) Error(line string, params ...interface{}) error {
	return f.Log(Error, line, params...)
}

// Critical logs a critical line
func (f *facility) Critical(line string, params ...interface{}) error {
	return f.Log(Critical, line, params...)
}

// Alert logs an alert line
func (f *facility) Alert(line string, params ...interface{}) error {
	return f.Log(Alert, line, params...)
}

// Emergency logs an emergency line
func (f *facility) Emergency(line string, params ...interface{}) error {
	return f.Log(Emergency, line, params...)
}
