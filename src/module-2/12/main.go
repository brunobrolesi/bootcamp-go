package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	LastName  string
	RG        string
	Admission time.Time
}

func (s Student) detail() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Sobrenome:", s.LastName)
	fmt.Println("RG:", s.RG)
	fmt.Println("Admission:", s.Admission.Format(time.RFC1123))
}

func main() {
	s1 := Student{
		Name:      "João",
		LastName:  "da Silva",
		RG:        "123456",
		Admission: time.Date(1980, 01, 15, 0, 0, 0, 0, time.UTC),
	}
	s1.detail()
	s2 := Student{
		Name:      "Maria",
		LastName:  "Aparecida",
		RG:        "654321",
		Admission: time.Now().UTC(),
	}
	s2.detail()
}
