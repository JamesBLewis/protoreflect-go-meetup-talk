package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"go-meetup-talk/out/example.com/project/protos/person"
)

func main() {
	personMessage := person.Person{
		Name:  "James",
		Id:    1,
		Email: "james.lewis2@anz.com",
	}
	data, err := proto.Marshal(&personMessage)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded Person: %v\n", data)
}
