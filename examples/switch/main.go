package main

import (
	_ "embed"
	"fmt"

	"github.com/golang/protobuf/proto"

	"go-meetup-talk/out/example.com/project/protos/person"
)

//go:embed person.bytes
var personData []byte

func main() {

	// Try to unmarshal into each known type
	msgA := &MessageA{}
	msgB := &person.Person{}

	if err := proto.Unmarshal(personData, msgA); err == nil {
		fmt.Println("The data represents a MessageA")
	} else if err := proto.Unmarshal(personData, msgB); err == nil {
		fmt.Println("The data represents a MessageB")
	} else {
		fmt.Println("The data does not represent a known type")
	}
}
