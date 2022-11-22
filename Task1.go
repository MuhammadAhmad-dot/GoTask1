package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// function to get struct as parameter and print its values.
func printvalues(obj1 Person) {
	fmt.Println("Name : ", obj1.name)
	fmt.Println("Age : ", obj1.age)
	fmt.Println("Job : ", obj1.job)
	fmt.Println("Salary : ", obj1.salary)
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Muhammad Ahmad"
	pers1.age = 23
	pers1.job = "Student"
	pers1.salary = 10000

	// Pers2 specification
	pers2.name = "Waqas Ahmad"
	pers2.age = 21
	pers2.job = "Marketing"
	pers2.salary = 5000

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("\n")
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	//Access and print Pers1 info by passing struct as function parameter
	fmt.Println("\nperson1 info\n")
	printvalues(pers1)
	//Access and print Pers2 info by passing struct as function parameter
	fmt.Println("\nperson2 info\n")
	printvalues(pers2)
}