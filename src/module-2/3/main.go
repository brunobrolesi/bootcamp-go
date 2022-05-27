package main

import "fmt"

func main() {
	word := "duckbill"
	fmt.Printf("Word length: %d \n", len(word))
	for _, letter := range word {
		fmt.Println(string(letter))
	}
}
