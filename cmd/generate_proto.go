package cmd

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func generateProto() error {
	protocExc, err := exec.LookPath("protoc")
	if err != nil {
		return &exec.Error{
			Name: "protoc executable not found",
			Err:  errors.New("protoc executable not found"),
		}
	}

	command := &exec.Cmd{
		Path: protocExc,
		Args: []string{"",
			"-I=samples", "--go_out=pb",
			protoFile},
		Env:          nil,
		Dir:          "",
		Stdin:        nil,
		Stdout:       nil,
		Stderr:       nil,
		ExtraFiles:   nil,
		SysProcAttr:  nil,
		Process:      nil,
		ProcessState: nil,
	}
	log.Info("proto generate command: %s", command.String())
	err = command.Run()
	if err != nil {
		return err
	}
	return nil
	// protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
}
