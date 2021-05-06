package cmd_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/evrenkutar/randevent/cmd"
)

func TestGenerateProto(t *testing.T) {
	fileContent := `syntax = "proto3";

package test_person;
option go_package = "github.com/evrenkutar/randevent/pb";

message PersonTest {
  string name = 1;
}
	`
	os.Chdir("..")
	err := ioutil.WriteFile("/tmp/test.proto", []byte(fileContent), 0655)
	if err != nil {
		panic("cannot write file")
	}
	err = cmd.GenerateProto("/tmp/test.proto")
	if err != nil {
		panic(err.Error())
	}

}
