# RANDEVENT

The purpose of this package is to generate events and push them to a kafka stream.
With this you can work with your event based application.
For now, it only covers proto buffers.

## Requirements

- protobuf 3.15.8

## build

`go build`

## run

#### 1. to generate proto files:
`randevent generate -p=samples/person.proto`

#### 2. to emit events
`randevent emit -i=100`
