package cmd

import (
	"fmt"
	"time"
)

type Emitter struct {
	tickChan chan time.Time
	doneChan chan bool
}

func (e *Emitter) emit() {
	for {
		select {
		case tick := <-e.tickChan:
			fmt.Printf("time is %v \n\n", tick)
		case e.doneChan <- true:
			return
		}
	}

}
