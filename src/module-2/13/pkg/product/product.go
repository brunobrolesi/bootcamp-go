package product

import "errors"

type Product interface {
	CalcCost() float64
}

const (
	Small  = "small"
	Medium = "medium"
	Big    = "big"
)

type product struct {
	Name  string
	Size  string
	Price float64
}

func (p product) CalcCost() float64 {
	if p.Size == Small {
		return p.Price
	}

	if p.Size == Medium {
		return p.Price * 1.03
	}

	return (p.Price * 1.06) + 2500
}

func NewProduct(name string, size string, price float64) (Product, error) {
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

func isValidSize(size string) bool {
	sizes := map[string]string{Small: Small, Medium: Medium, Big: Big}
	_, ok := sizes[size]

	if !ok {
		return false
	}
	return true
}
