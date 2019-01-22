package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/lajosbencz/glo"
)

func main() {
	log := glo.NewFacility()

	// write everything to a buffer
	bfr := bytes.NewBufferString("")
	handlerBfr := glo.NewHandler(bfr)
	log.PushHandler(handlerBfr)

	// write only errors and above using a short format
	handlerStd := glo.NewHandler(os.Stdout)
	formatter := glo.NewFormatter("%[2]s: %[3]s %[4]#")
	filter := glo.NewFilterLevel(glo.Error)
	handlerStd.SetFormatter(formatter)
	handlerStd.PushFilter(filter)
	log.PushHandler(handlerStd)

	fmt.Println("Log output:")
	fmt.Println(strings.Repeat("=", 70))
	log.Info("Only written to the buffer")
	log.Alert("Written to both buffer and stdout")

	fmt.Println("")
	fmt.Println("Buffer contents:")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println(bfr.String())
}
