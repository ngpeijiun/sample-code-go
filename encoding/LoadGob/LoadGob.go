package main

import (
	"bytes"
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

func (p Person) String() string {
	var buf bytes.Buffer

	buf.WriteString(p.Name.Personal)
	buf.WriteRune(' ')
	buf.WriteString(p.Name.Family)

	for _, v := range p.Emails {
		buf.WriteRune('\n')
		buf.WriteString(v.Kind)
		buf.WriteString(": ")
		buf.WriteString(v.Address)
	}

	return buf.String()
}

func main() {
	var p Person

	loadGob("person.gob", &p)

	fmt.Println("Person", p)
}

func loadGob(fileName string, value interface{}) {
	file, err := os.Open(fileName)

	checkError(err)

	defer file.Close()

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(value)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}
