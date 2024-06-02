package main

import (
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"

	"go-meetup-talk/out/example.com/project/protos/person"
)

const ProtoFileName = "person.bytes"
const JsonFileName = "person.json"

func main() {
	personMessage := &person.Person{
		Name:  "James",
		Id:    1,
		Email: "james.lewis2@anz.com",
	}
	writeProto(personMessage)
	writeJson(personMessage)
}

func writeProto(message proto.Message) {
	data, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}
	// Create a file
	file, err := os.Create(ProtoFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Proto data written to %s file successfully. len %v\n", ProtoFileName, len(data))
}

func writeJson(message proto.Message) {
	data, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	// Create a file
	file, err := os.Create(JsonFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Print("%v\n", data)

	// Write data to the file
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Json data written to %s file successfully. len %v\n", JsonFileName, len(data))
}
