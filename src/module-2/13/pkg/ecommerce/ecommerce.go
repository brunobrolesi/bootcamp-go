package ecommerce

import "github.com/brunobrolesi/bootcamp-go/src/module-2/13/pkg/product"

type Ecommerce interface {
	Total() float64
	Add(p product.Product) Ecommerce
}

type store struct {
	products []product.Product
}

func NewStore() Ecommerce {
	store := store{}
	return store
}

func (s store) Total() float64 {
	total := 0.0
	for _, product := range s.products {
		total += product.CalcCost()
	}
	return total
}

func (s store) Add(p product.Product) Ecommerce {
	s.products = append(s.products, p)
	return s
}
