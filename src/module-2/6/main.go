package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Printf("Benjamin age: %d \n", employees["Benjamin"])

	count := 0
	for _, value := range employees {
		if value > 21 {
			count++
		}
	}
	fmt.Printf("Employees over 21 years old: %d \n", count)

	employees["Frederico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
