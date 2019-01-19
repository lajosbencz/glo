package glo

// LoggerDebug logs a debug line
type LoggerDebug interface {
	Debug(line string, params ...interface{}) Logger
}

// LoggerInfo logs an info line
type LoggerInfo interface {
	Info(line string, params ...interface{}) Logger
}

// LoggerNotice logs a notice line
type LoggerNotice interface {
	Notice(line string, params ...interface{}) Logger
}

// LoggerWarning logs a warning line
type LoggerWarning interface {
	Warning(line string, params ...interface{}) Logger
}

// LoggerError logs an error line
type LoggerError interface {
	Error(line string, params ...interface{}) Logger
}

// LoggerCritical logs a critical line
type LoggerCritical interface {
	Critical(line string, params ...interface{}) Logger
}

// LoggerAlert logs an alert line
type LoggerAlert interface {
	Alert(line string, params ...interface{}) Logger
}

// LoggerEmergency logs an emergency line
type LoggerEmergency interface {
	Emergency(line string, params ...interface{}) Logger
}

// Debug logs a debug line
func (l *logger) Debug(line string, params ...interface{}) Logger {
	l.Log(Debug, line, params...)
	return l
}

// Info logs an info line
func (l *logger) Info(line string, params ...interface{}) Logger {
	l.Log(Info, line, params...)
	return l
}

// Notice logs a notice line
func (l *logger) Notice(line string, params ...interface{}) Logger {
	l.Log(Notice, line, params...)
	return l
}

// Warning logs a warning line
func (l *logger) Warning(line string, params ...interface{}) Logger {
	l.Log(Warning, line, params...)
	return l
}

// Error logs an error line
func (l *logger) Error(line string, params ...interface{}) Logger {
	l.Log(Error, line, params...)
	return l
}

// Critical logs a critical line
func (l *logger) Critical(line string, params ...interface{}) Logger {
	l.Log(Critical, line, params...)
	return l
}

// Alert logs an alert line
func (l *logger) Alert(line string, params ...interface{}) Logger {
	l.Log(Alert, line, params...)
	return l
}

// Emergency logs an emergency line
func (l *logger) Emergency(line string, params ...interface{}) Logger {
	l.Log(Emergency, line, params...)
	return l
}
