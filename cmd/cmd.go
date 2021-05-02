package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	"time"
)

var (
	protoFile string
	interval  int32
	rootCmd   = &cobra.Command{
		Use:   "randevent",
		Short: "Randevent is random event pusher to kafka",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("hello friend!")
		},
	}
	genProtoCmd = &cobra.Command{
		Use: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			err := generateProto()
			if err != nil {
				log.Fatalf("proto generation failed: %d", err.Error())
			}
			log.Info("proto files generated")
		},
	}
	emitCmd = &cobra.Command{
		Use: "emit",
		Run: func(cmd *cobra.Command, args []string) {
			emit()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	genProtoCmd.PersistentFlags().StringVarP(&protoFile, "protofile", "p", "", "proto file to generate random events")
	emitCmd.PersistentFlags().Int32VarP(&interval, "interval", "i", 50, "interval to generate")
	err := genProtoCmd.MarkPersistentFlagRequired("protofile")
	if err != nil {
		return
	}
	rootCmd.AddCommand(emitCmd, genProtoCmd)
}

func emit() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		log.Error("kafka producer can not be initialized")
		panic(err)
	}

	emitter := Emitter{
		doneChan:      make(chan bool),
		tickChan:      make(chan time.Time),
		kafkaProducer: kafkaProducer,
	}
	go emitter.emit()
	go start(done, ticker, emitter)

	time.Sleep(10 * time.Second)
	done <- true
}
