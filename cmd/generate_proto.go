package cmd

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"

	log "github.com/sirupsen/logrus"
)

func GenerateProto(f string) error {
	protocExc, err := exec.LookPath("protoc")
	if err != nil {
		return &exec.Error{
			Name: "protoc executable not found",
			Err:  err,
		}
	}

	d, err := ioutil.ReadFile(f)
	if err != nil {
		return &exec.Error{
			Name: "file cannot be read",
			Err:  err,
		}
	}
	_, protoFileName := path.Split(f)
	err = ioutil.WriteFile(fmt.Sprintf("samples/%s", protoFileName), d, 0655)
	if err != nil {
		return &exec.Error{
			Name: "file cannot be saved",
			Err:  err,
		}
	}

	command := &exec.Cmd{
		Path: protocExc,
		Args: []string{"",
			"-I=samples", "--go_out=pb",
			fmt.Sprintf("samples/%s", protoFileName)},
	}
	log.Infof("proto generate command: %s", command.String())
	err = command.Run()
	if err != nil {
		return err
	}
	return nil
	// protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
}
