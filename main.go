package main

import (
	"log"

	pb "github.com/micnncim/protoc-gen-validate-playground/example"
)

func main() {
	p := &pb.Person{
		Id:    1000,
		Email: "person@mail.net",
		Name:  "Protocol Buffers",
		Home: &pb.Person_Location{
			Lat: 100,
			Lng: 0,
		},
	}
	err := p.Validate()
	if e, ok := err.(pb.PersonValidationError); ok {
		log.Println(e.Error())
		log.Println(e.ErrorName())
		log.Println(e.Field())
		log.Println(e.Reason())
		log.Println(e.Cause())
		log.Println(e.Key())
	}
}
