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
	log.PushHandler(glo.NewHandler(os.Stdout).SetFormatter(glo.NewFormatter("[%[1]s] %[2]s %[3]v")))

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		for i := 0; i < loopCount; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleepMax-sleepMin)+sleepMin))
			log.Log(glo.Debug, "Hello Log!", map[string]string{"x": "foo2", "y": "bar"})
			log.Log(glo.Debug, "Detailed info to debug")
		}
		waitGrp.Done()
	}()

	go func() {
		for i := 0; i < loopCount; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleepMax-sleepMin)+sleepMin))
			log.Log(glo.Info, "Stream this!")
		}
		waitGrp.Done()
	}()

	waitGrp.Wait()
}
