package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	dfunc, _ := Animal(dog)
	fmt.Printf("quandidade de alimento do(s) cachorro(s) (em gramas): %d gramas\n", dfunc(5))
	cfunc, _ := Animal(cat)
	fmt.Printf("quandidade de alimento do(s) gato(s) (em gramas): %d gramas\n", cfunc(8))
	_, err := Animal("invalid animal")
	if err != nil {
		fmt.Println("animal inválido")
	}
}

func dogFunc(quantity int) int {
	return quantity * 10000
}

func catFunc(quantity int) int {
	return quantity * 5000
}

func hamsterFunc(quantity int) int {
	return quantity * 250
}

func tarantulaFunc(quantity int) int {
	return quantity * 150
}

func Animal(t string) (func(quantity int) int, error) {
	animalFunctions := map[string]func(quantity int) int{
		dog:       dogFunc,
		cat:       catFunc,
		hamster:   hamsterFunc,
		tarantula: tarantulaFunc,
	}

	selected := animalFunctions[t]
	if selected == nil {
		return nil, errors.New("invalid animal type")
	}

	return selected, nil
}
