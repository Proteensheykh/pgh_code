package main

import (
	"fmt"
)

func main() {
	age := -71

	var price, category = getTicketInfo(age)

	fmt.Printf("The ticket price for %v is N%v.", category, price)

}

func getTicketInfo(age int) (int, string) {
	if age > 0 && age <= 12 {
		return 5, "children"
	} else if age > 12 && age <= 64 {
		return 10, "Adult"
	} else if age >= 65 {
		return 7, "Senior"
	} else {
		return 0, "not allowed"
	}
}
