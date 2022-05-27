package main

import "fmt"

type Person struct {
	Age          int
	IsEmployed   bool
	TimeEmployed int
	Salary       float32
}

func main() {
	p1 := Person{
		Age:          24,
		IsEmployed:   true,
		TimeEmployed: 2,
		Salary:       102000,
	}
	p2 := Person{
		Age:          20,
		IsEmployed:   false,
		TimeEmployed: 0,
		Salary:       0,
	}
	p3 := Person{
		Age:          25,
		IsEmployed:   true,
		TimeEmployed: 3,
		Salary:       80000,
	}

	getLoan(p1)
	getLoan(p2)
	getLoan(p3)
}

func canReceiveLoan(p Person) bool {
	if p.Age < 22 || !p.IsEmployed || p.TimeEmployed < 1 {
		return false
	}

	return true
}

func getLoan(p Person) {
	if !canReceiveLoan(p) {
		fmt.Println("The client is not qualified to receive de loan")
		return
	}

	if p.Salary < 100000 {
		fmt.Println("The client can receive the loan")
		return
	}

	fmt.Println("The client can receive the loan without fees")
}
