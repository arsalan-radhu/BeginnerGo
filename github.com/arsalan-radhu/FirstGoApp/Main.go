package main

import (
	"fmt"
	"strconv"
)

func newExampleFunction(firstNum int, secondNum int) int {
	multipliedNumber := firstNum * secondNum
	return multipliedNumber
}

func main() {

	//Variable Declaration
	// var i int
	// i =42

	//var i int = 42 --> good for telling which data type is used

	// Code asssumes the data type
	i := 42
	fmt.Println(i)

	var j string = strconv.Itoa(i)
	fmt.Printf("%v,%T\n", j, j)

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekdays := days[1:5]
	fmt.Println(weekdays)

	students := map[string]string{
		"Arsalan": "Radhu",
		"Apoorv":  "Gupta",
	}

	fmt.Println(students["Arsalan"])

	type Person struct {
		name    string
		age     int
		address string
	}

	newPerson := Person{name: "Arsalan", age: 20, address: "7 Bradford"}
	newPeople := [...]Person{{name: "AR", age: 22, address: "123abc"}, {name: "BC", age: 24, address: "asda221"}}
	newPerson.age = 21

	var person2 Person = newPeople[1]
	fmt.Println(person2.name)
	fmt.Println(newPeople[0])

	fmt.Println("Name: " + newPerson.name + "\nAge: " + strconv.Itoa(newPerson.age) + "\nAddress: " + newPerson.address)
	type Team struct {
		name    string
		players [5]Person
	}

	players := [...]Person{{name: "Cristiano Ronaldo"}, {name: "Bruno Fernandes"}, {name: "De Gea"}, {name: "Jadon Sancho"}, {name: "Marcus Rashford"}}

	mufc := Team{name: "Manchester United", players: players}

	fmt.Println(mufc)

	fmt.Println("Enter First Number: ")
	var fnum int
	fmt.Scanln(&fnum)

	var snum int
	fmt.Println("Enter Second Number: ")
	fmt.Scanln(&snum)
	multiplication := newExampleFunction(fnum, snum)

	fmt.Println("The multiplication of " + strconv.Itoa(fnum) + " and " + strconv.Itoa(snum) + " is " + strconv.Itoa(multiplication))

}
