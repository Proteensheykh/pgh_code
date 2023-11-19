package main

import (
	"fmt"
)

func main() {
	year := 2000

	if isLeapYear(year) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
