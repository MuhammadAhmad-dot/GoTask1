package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string

	employees []employee
}

func main() {
	var emp1 employee
	var emp2 employee
	var emp3 employee

	//emp1 specification
	emp1.name = "Muhammad Ahmad"
	emp1.salary = 10000
	emp1.position = "Software Developer"

	//emp2 specification
	emp2.name = "Waqas Ahmad"
	emp2.salary = 20000
	emp2.position = "Web Developer"

	//emp3 specification
	emp3.name = "Hanzla Sibghat"
	emp3.salary = 15000
	emp3.position = "Data Analyst"

	emplys := []employee{emp1, emp2, emp3}

	var comp company

	//comp1 spesification
	comp.companyName = "Tesla"
	comp.employees = emplys

	fmt.Println("\n")
	fmt.Println("Company Name : ", comp.companyName)
	fmt.Println("Employees : ", comp.employees)

}