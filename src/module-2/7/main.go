package main

import "fmt"

func main() {
	fmt.Printf("Tax of up to R$50.000: %.2f \n", calcTaxes(50000))
	fmt.Printf("Tax of up to R$150.000: %.2f \n", calcTaxes(150000))
	fmt.Printf("Tax more than R$150.000: %.2f \n", calcTaxes(150001))
}

func calcTaxes(salary float64) float64 {
	if salary <= 50000.00 {
		return 0.0
	} else if salary <= 150000.00 {
		return salary * 0.17
	}
	return salary * 0.27
}
