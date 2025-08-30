package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	employeeId string
}

func (e Employee) printInfo() {
	fmt.Printf("name %s \n age %d \n employid %s\n", e.name, e.age, e.employeeId)
}

func main() {
	emp := Employee{
		Person: Person{
			name: "test",
			age:  24,
		},
		employeeId: "33",
	}

	emp.printInfo()
}
