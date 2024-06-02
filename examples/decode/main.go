package main

import (
	"io"
	"os"

	"google.golang.org/protobuf/proto"

	"go-meetup-talk/out/example.com/project/protos/person"
)

const FileName = "person.bytes"

func main() {
	var personMessage person.Person

	// open file
	file, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read file
	byteData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Unmarshal back into proto message
	err = proto.Unmarshal(byteData, &personMessage)
	if err != nil {
		panic(err)
	}
}
