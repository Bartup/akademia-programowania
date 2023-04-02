package main

import (
	"github.com/google/uuid"
	"github.com/grupawp/appdispatcher"
	"log"
)

type Student struct {
	FirstName     string
	LastName      string
	applicationID uuid.UUID
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func (s Student) ApplicationID() string {
	return s.applicationID.String()
}

func main() {
	stud := Student{
		FirstName:     "Bartosz",
		LastName:      "Grams",
		applicationID: uuid.New(),
	}

	code, err := appdispatcher.Submit(stud)

	log.Println(code)
	log.Println(err)
}
