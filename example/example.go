package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/lajosbencz/glo"
)

const (
	loopCount int = 50
	sleepMin  int = 1
	sleepMax  int = 50
)

func main() {
	taskRemain := 2
	waitChn := make(chan bool, taskRemain)

	log := glo.NewLogger()
	log.AddHandler(glo.NewHandler(os.Stdout, glo.NewFormatter(glo.DefaultFormat)))
	log.AddHandler(glo.NewHandler(os.Stdout, glo.NewFormatter("[%[1]s] %[2]s %[3]v")))

	go func() {
		for i := 0; i < loopCount; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleepMax-sleepMin)+sleepMin))
			log.Log(glo.Debug, "Hello Log!", map[string]string{"x": "foo2", "y": "bar"})
			log.Log(glo.Debug, "Detailed info to debug")
		}
		waitChn <- true
	}()

	go func() {
		for i := 0; i < loopCount; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleepMax-sleepMin)+sleepMin))
			log.Log(glo.Info, "Stream this!")
		}
		waitChn <- true
	}()

	for {
		select {
		case <-waitChn:
			taskRemain--
		}
		if taskRemain < 1 {
			break
		}
	}
}
