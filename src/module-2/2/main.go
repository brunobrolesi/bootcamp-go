package main

import "fmt"

func main() {
	var temperature float32 = 16.4
	var humidity uint8 = 75
	var pressure float32 = 1018

	fmt.Printf("Temperature: %.2f ºC \n", temperature)
	fmt.Printf("Humidity: %d %% \n", humidity)
	fmt.Printf("Pressure: %.2f hPa \n", pressure)
}
