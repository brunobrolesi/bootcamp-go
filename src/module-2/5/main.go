package main

import "fmt"

func main() {
	n1, n2, n3 := 1, 5, 12

	fmt.Println(getMonth(n1))
	fmt.Println(getMonth(n2))
	fmt.Println(getMonth(n3))
}

func getMonth(n int) string {
	monthsMap := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	return monthsMap[n]
}
