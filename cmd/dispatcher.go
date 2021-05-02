package cmd

import "errors"

type Protobuf interface {
}

type Avro interface {
	// to be implemented
}

type Dispatcher struct {
	p          Protobuf
	a          Avro
	Serializer interface{}
}

func (d *Dispatcher) Dispatch(input string) (string, error) {
	if input == "protobuf" {
		d.Serializer = d.p
	} else if input == "avro" {
		d.Serializer = d.a
	} else {
		return "", errors.New("unknown serializer")
	}
	return "", nil
}
