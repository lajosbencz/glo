package glo

import "errors"

// ErrValidatorLevelRange invalid arguments
var ErrValidatorLevelRange = errors.New("min cannot be greater than max")

// Validator checks a log line against custom logic
type Validator interface {
	Validate(Level, string, ...interface{}) bool
}

// NewValidatorLevel checks the level of the log (inclusive)
func NewValidatorLevel(min Level) Validator {
	return &validatorLevel{min}
}

// NewValidatorLevelRange checks if the level of the log is in a range (inclusive)
func NewValidatorLevelRange(min, max Level) Validator {
	if min > max {
		t := max
		max = min
		min = t
	}
	return &validatorLevelRange{min, max}
}

type validatorLevel struct {
	level Level
}

func (v *validatorLevel) Validate(level Level, line string, params ...interface{}) bool {
	return level >= v.level
}

type validatorLevelRange struct {
	min Level
	max Level
}

func (v *validatorLevelRange) Validate(level Level, line string, params ...interface{}) bool {
	return level >= v.min && level <= v.max
}
