package main

import (
	"fmt"
	"io/ioutil"
	"log"

	simple "CodeGenGO/Example.Simple/simplepb/simple"
	"CodeGenGO/enumeration/enumpb"
	"CodeGenGO/example/complexMessage"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	//Read write Demo with struct from generated proto buf
	ReadAndWriteDemo(sm)

	// Json converstion Demo with struct  from generated proto buf
	JsonDemo(sm)

	// Enum with struct from generated proto buf
	doEnum()

	// Complex message handling with struct from generated proto buf
	doComplex()

}

func doComplex() {
	cm := complexMessage.ComplexMessage{
		OneDummy: &complexMessage.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexMessage.DummyMessage{
			{
				Id:   2,
				Name: "Second Message",
			},
			{
				Id:   3,
				Name: "Third Message",
			},
		},
	}

	fmt.Println("data inside complex Message", cm)
}

func doEnum() {

	em := enumpb.EnumMessage{
		Id:        42,
		DayOfWeek: enumpb.DayofWeek_MONDAY,
	}

	em.DayOfWeek = enumpb.DayofWeek_FRIDAY

	fmt.Println("Data from enumeration", em)
}

func ReadAndWriteDemo(sm proto.Message) {
	writeToFile("MyCustomProtoFile.bin", sm)

	// Create a pointer to the struct and pass the empty struct to the read file function to get it populated

	sm2 := doSimple()
	readFile("MyCustomProtoFile.bin", sm2)

	fmt.Println("data inside sm2 ", sm2)
}

func JsonDemo(sm2 proto.Message) {
	data := toJson(sm2)
	fmt.Println("Data in JSon", data)

	sm3 := &simple.SimpleMessage{}
	fromJson(data, sm3)
	fmt.Println("Data from unmarshelling ", sm3)
}

func readFile(fname string, pb proto.Message) error {
	out, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong while reading file", err)
		return err
	}

	if err := proto.Unmarshal(out, pb); err != nil {
		log.Fatalln("Couldnt pass data from file to content read", err)
		return err
	}

	return nil

}

func toJson(pb proto.Message) string {
	marsheller := jsonpb.Marshaler{}
	str, err := marsheller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cant convert to json")
		return ""
	}
	return str
}

func fromJson(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldnt unmarshal json")
	}
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cant serialize data to bytes")
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0666); err != nil {
		log.Fatalln("Cant write to file")
		return err
	}

	fmt.Println("Data has been written")

	return nil
}

func doSimple() *simple.SimpleMessage {

	m := simple.SimpleMessage{
		Id:         333,
		IsSimple:   true,
		Name:       "MySimpleMessage",
		SampleList: []int32{1, 4, 6, 7},
	}

	fmt.Println(m.Id, m.IsSimple, m.Name, m.SampleList)

	m.Name = "Hitesh"

	fmt.Println(m)

	fmt.Println(m.GetName())

	return &m
}
