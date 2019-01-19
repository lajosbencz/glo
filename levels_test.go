package glo

import (
	"testing"
)

func TestLevelSnap(t *testing.T) {
	if levelSnap(Alert) != Alert {
		t.Error("levelSnap equal")
	} else if levelSnap(Alert-1) != Critical {
		t.Error("levelSnap lesser")
	} else if levelSnap(Alert+1) != Alert {
		t.Error("levelSnap greater")
	}
}

func TestLevelString(t *testing.T) {
	for lvl, exp := range levelNames {
		got := lvl.String()
		if got != exp {
			t.Errorf("invalid toString format for Level %d, expected(%s) got(%s)", lvl, exp, got)
		}
	}
}

func TestLevelOrder(t *testing.T) {
	testLevelGreater(t, Emergency, Alert)
	testLevelGreater(t, Alert, Critical)
	testLevelGreater(t, Critical, Error)
	testLevelGreater(t, Error, Warning)
	testLevelGreater(t, Warning, Notice)
	testLevelGreater(t, Notice, Info)
	testLevelGreater(t, Info, Debug)
}

func testLevelGreater(t *testing.T, a, b Level) {
	if a <= b {
		t.Errorf("%[2]s(%[2]d) is greater than %[1]s(%[1]d)", a, b)
	}
}
