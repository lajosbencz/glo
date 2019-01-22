package main

import "github.com/lajosbencz/glo"

func main() {
	// Info - Warning will go to os.Stdout
	// Error - Emergency will go to os.Stderr
	log := glo.NewStdFacility()

	// goes to os.Stdout
	log.Debug("Detailed debug line: %#v", map[string]string{"x": "foo", "y": "bar"})

	// goes to os.Stderr
	log.Error("Oooof!")
}
