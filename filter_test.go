package glo

import "testing"

func TestFilterLevel(t *testing.T) {
	v := NewFilterLevel(Warning)
	if v.Check(Debug, "") {
		t.Errorf("%s is lower than %s", Debug, Warning)
	}
	if !v.Check(Critical, "") {
		t.Errorf("%s is higher than %s", Critical, Warning)
	}
}

func TestFilterLevelRange(t *testing.T) {
	v := NewFilterLevelRange(Warning, Info)
	if v.Check(Debug, "") {
		t.Errorf("%s should be invalid", Debug)
	}
	if !v.Check(Notice, "") {
		t.Errorf("%s should be valid", Error)
	}
	if v.Check(Error, "") {
		t.Errorf("%s should be invalid", Error)
	}
}
