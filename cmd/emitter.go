package cmd

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/evrenkutar/randevent/pb/github.com/evrenkutar/randevent/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Emitter struct {
	tickChan chan time.Time
	doneChan chan bool
}

func (e *Emitter) emit() {
	for {
		select {
		case <-e.tickChan:
			event := generateNewEvent()
			fmt.Printf("event is %v \n\n", event) // todo: emit this to kafka
		case e.doneChan <- true:
			return
		}
	}

}

func generateNewEvent() string {
	p := pb.PersonCreated{
		Uuid:      gofakeit.UUID(),
		Firstname: gofakeit.Name(),
		Lastname:  gofakeit.LastName(),
		CreatedAt: timestamppb.Now(),
	}
	return p.String()
}
