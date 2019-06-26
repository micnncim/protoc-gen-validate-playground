package main

import (
	"log"

	pb "github.com/micnncim/protoc-gen-validate-playground/example"
)

func main() {
	p := new(pb.Person)
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
