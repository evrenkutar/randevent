package cmd

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/evrenkutar/randevent/pb/github.com/evrenkutar/randevent/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"time"
)

type Emitter struct {
	tickChan      chan time.Time
	doneChan      chan bool
	kafkaProducer *kafka.Producer
}

func (e *Emitter) emit() {
	go func() {
		for e := range e.kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	for {
		select {
		case <-e.tickChan:
			event := generateNewEvent()
		    topic := "testTopic"
			e.kafkaProducer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value: []byte(event),
			}, nil)
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
