package cmd

import (
	"fmt"
	"time"
)

type Emitter struct {
	tickChan chan time.Time
	doneChan chan bool
	event    string
}

func (e *Emitter) emit() {
	for {
		select {
		case <-e.tickChan:
			fmt.Printf("event is %v \n\n", e.event)
		case e.doneChan <- true:
			return
		}
	}

}
