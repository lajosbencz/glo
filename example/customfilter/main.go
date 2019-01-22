package main

import (
	"os"
	"regexp"

	"github.com/lajosbencz/glo"
)

func main() {
	handler := glo.NewHandler(os.Stdout)
	filterEmptyLines := &filterRgx{regexp.MustCompile(`^.+$`)}
	handler.PushFilter(filterEmptyLines)

	log := glo.NewFacility()
	log.PushHandler(handler)

	log.Debug("", "format is empty, should be ignored")
	log.Debug("only this should appear at the output")
}

type filterRgx struct {
	rgx *regexp.Regexp
}

func (f *filterRgx) Check(level glo.Level, line string, params ...interface{}) bool {
	return f.rgx.MatchString(line)
}
