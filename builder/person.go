package main

import "fmt"

type Person struct {
	// Address
	StreetAddress, Postcode, City string

	// Job
	Company, Position string
	AnnualIncome      int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (p *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	p.person.StreetAddress = streetAddress
	return p
}
func (p *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	p.person.City = city
	return p
}
func (p *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	p.person.Postcode = postcode
	return p
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (j *PersonJobBuilder) WorksAt(company string) *PersonJobBuilder {
	j.person.Company = company
	return j
}
func (j *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	j.person.Position = position
	return j
}
func (j *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	j.person.AnnualIncome = annualIncome
	return j
}

func (j *PersonBuilder) Build() *Person {
	return j.person
}

func main() {
	pb := NewPersonBuilder()
	pb.Lives().
		At("123 London Road").
		In("London").
		WithPostCode("SW12BC").
	Works().
		WorksAt("Fabrikam").
		AsA("Engineer").
		Earning(123000)
	person := pb.Build()
	fmt.Println(person)
}