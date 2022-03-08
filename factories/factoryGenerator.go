package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			name,
			position,
			annualIncome,
		}
	}
}

func main() {
	managerEmployee := NewEmployeeFactory("Manager", 120000)
	staffEmployee := NewEmployeeFactory("Staff", 75000)

	manager := managerEmployee("John")
	staff := staffEmployee("Seun")

	fmt.Println(manager)
	fmt.Println(staff)
}