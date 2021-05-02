package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	protoFile string
	interval  int32
	rootCmd   = &cobra.Command{
		Use:   "randevent",
		Short: "Randevent is random event pusher to kafka",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Run()
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
	rootCmd.PersistentFlags().StringVarP(&protoFile, "protofile", "p", "", "proto file to generate random events")
	rootCmd.PersistentFlags().Int32VarP(&interval, "interval", "i", 50, "interval to generate")
	err := rootCmd.MarkPersistentFlagRequired("protofile")
	if err != nil {
		return
	}
	err = rootCmd.MarkPersistentFlagRequired("interval")
	if err != nil {
		return
	}
}

func Run() {
	err := generateProto()
	if err != nil {
		log.Fatalf("proto generation failed: %d", err.Error())
	}
	log.Info("proto files generated")

	done := make(chan bool)
	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	emitter := Emitter{
		doneChan: make(chan bool),
		tickChan: make(chan time.Time),
		event:    "heyho",
	}
	go emitter.emit()
	go start(done, ticker, emitter)

	time.Sleep(10 * time.Second)
	done <- true
}
