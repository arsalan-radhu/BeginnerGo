package main

import (
	"fmt"
)

type Employee struct {
	Name string
}

func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println("Updated Employee Name: " + e.Name)
}

func main() {
	var employee Employee

	employee.Name = "Arsalan"
	fmt.Println("Old Employee Name: " + employee.Name)
	employee.UpdateName("Radhu")
	employee.PrintName()
}
