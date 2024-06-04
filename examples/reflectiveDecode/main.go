package main

import (
	_ "embed"
	"fmt"

	cepb "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	event "github.com/cloudevents/sdk-go/v2/event"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"

	"go-meetup-talk/out/example.com/project/protos/person"
)

func main() {
	// send
	PersonEvent := sendCloudEvent()

	// receive
	person := decodeCloudEvent(PersonEvent)
	prettyPrint(person)
}

func sendCloudEvent() event.Event {
	// send
	personMessage := person.Person{
		Name:  "James",
		Id:    1,
		Email: "james.lewis2@anz.com",
	}

	// Now fetch event metadata from its descriptor (with cache).
	desc := personMessage.ProtoReflect().Descriptor()
	cedataschema := desc.FullName()

	byteData, err := proto.Marshal(&personMessage)
	if err != nil {
		panic(err)
	}
	cloudevent := event.New(cloudevents.VersionV1)

	err = cloudevent.SetData(cepb.ContentTypeProtobuf, byteData)
	if err != nil {
		panic(err)
	}
	cloudevent.SetDataSchema(fmt.Sprintf("%s", cedataschema))
	return cloudevent
}

func decodeCloudEvent(eventData event.Event) proto.Message {
	dataSchema := eventData.DataSchema()
	byteData := eventData.Data()

	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(dataSchema))
	if err != nil {
		panic(err)
	}
	messageDesc, isMsg := desc.(protoreflect.MessageDescriptor)
	if !isMsg {
		panic("Not a message")
	}

	newMsg := dynamicpb.NewMessage(messageDesc)
	err = proto.Unmarshal(byteData, newMsg)
	if err != nil {
		panic(err)
	}
	return newMsg
}

func prettyPrint(msg proto.Message) {
	msg.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fmt.Printf("Field: %s, Value: %v\n", fd.Name(), v.Interface())
		return true
	})
}
