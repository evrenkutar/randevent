# RANDEVENT

The purpose of this package is to generate events and push them to a kafka stream.
With this you can work with your event based application.
For now, it only covers proto buffers.

## Requirements

- protobuf 3.15.8

## build

`go build`

## run

#### 1. generate proto files:
`randevent generate -p=samples/person.proto`

#### 2. start kafka in docker container
This step is produced following this [guide](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html#step-1-download-and-start-cp-using-docker).
`docker-compose up`

#### 3. emit events
`randevent emit -i=100`
