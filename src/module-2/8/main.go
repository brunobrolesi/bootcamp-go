package main

import (
	"errors"
	"fmt"
)

func main() {
	s1, err1 := calcAverage(7, 8, 7)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("Average: %.2f \n", s1)
	}

	s2, err2 := calcAverage(-7, 8, 7)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Average: %.2f \n", s2)
	}
}

func calcAverage(grades ...float64) (float64, error) {
	sum := 0.0
	for _, n := range grades {
		if n < 0 {
			return 0, errors.New("invalid grade provided")
		}
		sum += n
	}
	return sum / float64(len(grades)), nil
}
