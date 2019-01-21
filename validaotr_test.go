package glo

import "testing"

func TestValidatorLevel(t *testing.T) {
	v := NewValidatorLevel(Warning)
	if v.Validate(Debug, "") {
		t.Errorf("%s is lower than %s", Debug, Warning)
	}
	if !v.Validate(Critical, "") {
		t.Errorf("%s is higher than %s", Critical, Warning)
	}
}

func TestValidatorLevelRange(t *testing.T) {
	v := NewValidatorLevelRange(Warning, Info)
	if v.Validate(Debug, "") {
		t.Errorf("%s should be invalid", Debug)
	}
	if !v.Validate(Notice, "") {
		t.Errorf("%s should be valid", Error)
	}
	if v.Validate(Error, "") {
		t.Errorf("%s should be invalid", Error)
	}
}
