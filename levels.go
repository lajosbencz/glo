package glo

// Level defines the severity
type Level uint16

// String formats the Level
func (l Level) String() string {
	l = levelSnap(l)
	return levelNames[l]
}

const (
	// Debug severity
	Debug Level = 100
	// Info severity
	Info = 200
	// Notice severity
	Notice = 250
	// Warning severity
	Warning = 300
	// Error severity
	Error = 400
	// Critical severity
	Critical = 500
	// Alert severity
	Alert = 550
	// Emergency severity
	Emergency = 600
)

// LevelList incrementally lists severities
var LevelList []Level = []Level{
	Debug,
	Info,
	Notice,
	Warning,
	Error,
	Critical,
	Alert,
	Emergency,
}

// LevelNames maps a Level to a string
var levelNames = map[Level]string{
	Debug:     "DEBUG",
	Info:      "INFO",
	Notice:    "NOTICE",
	Warning:   "WARNING",
	Error:     "ERROR",
	Critical:  "CRITICAL",
	Alert:     "ALERT",
	Emergency: "EMERGENCY",
}

func levelSnap(in Level) Level {
	if in >= Emergency {
		return Emergency
	}
	if in >= Alert {
		return Alert
	}
	if in >= Critical {
		return Critical
	}
	if in >= Error {
		return Error
	}
	if in >= Warning {
		return Warning
	}
	if in >= Notice {
		return Notice
	}
	if in >= Info {
		return Info
	}
	return Debug
}
