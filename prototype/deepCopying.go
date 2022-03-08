package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}
type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(p)
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	d.Decode(&result)
	return &result
}
func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}, []string{}}

	jane := john
	// Deep Copy
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}
	jane.Name = "Jane"
	jane.Address.City = "Paris"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
