package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Product struct {
	Id       int
	Price    float64
	Quantity int
}

func main() {
	file := Init()
	defer file.Close()

	p1 := Product{
		Id:       1,
		Price:    15.99,
		Quantity: 2,
	}
	p2 := Product{
		Id:       2,
		Price:    5.99,
		Quantity: 4,
	}
	p3 := Product{
		Id:       3,
		Price:    150.99,
		Quantity: 1,
	}

	SaveProducts(file, p1)
	SaveProducts(file, p2)
	SaveProducts(file, p3)
}

func SaveProducts(f *os.File, p Product) {
	s := fmt.Sprintf("%d;%.2f;%d;", p.Id, p.Price, p.Quantity)
	_, err := io.WriteString(f, s)
	if err != nil {
		panic(err)
	}
}

func Init() *os.File {
	file, err := os.Create("./products.txt")

	if err != nil {
		panic(err)
	}

	return file
}

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
