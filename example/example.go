package main

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/lajosbencz/glo"
)

const (
	loopCount int = 50
	sleepMin  int = 1
	sleepMax  int = 50
)

func main() {

	log := glo.NewFacility()
	log.PushHandler(glo.NewHandler(os.Stdout))
	hShortInfo := glo.NewHandler(os.Stdout).PushValidator(glo.NewValidatorLevel(glo.Info)).SetFormatter(glo.NewFormatter("%[2]s: %[3]s %[4]v"))
	log.PushHandler(hShortInfo)

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		for i := 0; i < loopCount; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleepMax-sleepMin)+sleepMin))
			log.Log(glo.Info, "Hello Log!")
			log.Log(glo.Notice, "Detailed info to debug", map[string]string{"x": "foo", "y": "bar"})
		}
		waitGrp.Done()
	}()

	go func() {
		for i := 0; i < loopCount; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleepMax-sleepMin)+sleepMin))
			log.Log(glo.Error, "Ooof!")
			log.Debug("Only shows up in one handler")
		}
		waitGrp.Done()
	}()

	waitGrp.Wait()
}
