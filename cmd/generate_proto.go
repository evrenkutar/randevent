package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"

	log "github.com/sirupsen/logrus"
)

func generateProto() error {
	protocExc, err := exec.LookPath("protoc")
	if err != nil {
		return &exec.Error{
			Name: "protoc executable not found",
			Err:  errors.New("protoc executable not found"),
		}
	}

	d, err := ioutil.ReadFile(protoFile)
	if err != nil {
		return &exec.Error{
			Name: "file cannot be read",
			Err: errors.New("file cannot be read"),
		}
	}
	_, protoFileName := path.Split(protoFile)
	err = ioutil.WriteFile(fmt.Sprintf("samples/%s", protoFileName), d, 0655)
    if err != nil {
    	return &exec.Error{
			Name: "file cannot be saved",
			Err: errors.New("file cannot be saved"),
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
