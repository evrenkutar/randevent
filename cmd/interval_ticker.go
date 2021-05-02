package cmd

import (
	"time"
)

func start(done <-chan bool, ticker *time.Ticker, emitter Emitter) {
	for {
		select {
		case <-done:
			emitter.doneChan <- true
			return
		case tick := <-ticker.C:
			emitter.tickChan <- tick
		}
	}
}
