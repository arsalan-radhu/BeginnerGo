package main

import "fmt"

func main() {
	students := []string{"Arsalan", "Ammar", "Arif", "Viqar", "Summi", "Mariam"}

	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])

	}

	// There is no concept of while loops

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
