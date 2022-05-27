package main

import (
	"fmt"

	"github.com/brunobrolesi/bootcamp-go/src/module-2/13/pkg/ecommerce"
	"github.com/brunobrolesi/bootcamp-go/src/module-2/13/pkg/product"
)

func main() {
	p1, _ := product.NewProduct("Table", product.Medium, 820.0)
	p2, _ := product.NewProduct("Alexa", product.Small, 200.0)
	p3, _ := product.NewProduct("Refrigerator", product.Big, 2500.0)

	store := ecommerce.NewStore()

	store = store.Add(p1)
	fmt.Printf("Total Value: R$ %.2f \n", store.Total())
	store = store.Add(p2)
	fmt.Printf("Total Value: R$ %.2f \n", store.Total())
	store = store.Add(p3)
	fmt.Printf("Total Value: R$ %.2f \n", store.Total())

}
