package main

import (
	"errors"
	"fmt"
)

type Product interface {
	calcCost() float64
}

const (
	small  = "small"
	medium = "medium"
	big    = "big"
)

type product struct {
	Name  string
	Size  string
	Price float64
}

func (p product) calcCost() float64 {
	if p.Size == small {
		return p.Price
	}

	if p.Size == medium {
		return p.Price * 1.03
	}

	return (p.Price * 1.06) + 2500
}

func newProduct(name string, size string, price float64) (Product, error) {
	if !isValidSize(size) {
		return nil, errors.New("the size is invalid")
	}

	product := product{
		Name:  name,
		Size:  size,
		Price: price,
	}

	return product, nil
}

func newStore() Ecommerce {
	store := store{}
	return store
}

func isValidSize(size string) bool {
	sizes := map[string]string{small: small, medium: medium, big: big}
	_, ok := sizes[size]

	if !ok {
		return false
	}
	return true
}

type Ecommerce interface {
	Total() float64
	Add(p Product) Ecommerce
}

type store struct {
	products []Product
}

func (s store) Total() float64 {
	total := 0.0
	for _, product := range s.products {
		total += product.calcCost()
	}
	return total
}

func (s store) Add(p Product) Ecommerce {
	s.products = append(s.products, p)
	return s
}

func main() {
	p1, _ := newProduct("Table", medium, 820.0)
	p2, _ := newProduct("Alexa", small, 200.0)
	p3, _ := newProduct("Refrigerator", big, 2500.0)

	store := newStore()

	store = store.Add(p1)
	fmt.Printf("Total Value: R$ %.2f \n", store.Total())
	store = store.Add(p2)
	fmt.Printf("Total Value: R$ %.2f \n", store.Total())
	store = store.Add(p3)
	fmt.Printf("Total Value: R$ %.2f \n", store.Total())

}
