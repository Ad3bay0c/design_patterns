package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}
}

type person struct {
	name string
	age  int
}
type Person1 interface {
	SayHello()
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("My name is %s and I'm %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired\n")
}

func NewPerson1(name string, age int) Person1 {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name: name, age: age}
}

func main() {
	p := NewPerson1("James", 132)
	p.SayHello()
}
