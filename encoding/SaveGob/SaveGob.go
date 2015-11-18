package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name Name
	Emails []Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}

func main() {
	p := Person{
		Name: Name{ Family: "Newmarch", Personal: "Jan" },
		Emails: []Email{
			Email{ Kind: "home", Address: "jan@newmarch.name" },
			Email{ Kind: "work", Address: "j.newmarch@boxhill.edu.au" }}}

	saveGob("person.gob", p)
}

func saveGob(fileName string, value interface{}) {
	file, err := os.Create(fileName)

	checkError(err)

	defer file.Close()

	encoder := gob.NewEncoder(file)

	err = encoder.Encode(value)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
