package main

import (
	"fmt"
)

func main() {
	score := 23

	fmt.Println("The score is", score, "and the resulting grade is", getGrade(score))
}

func getGrade(score int) string {
	if score >= 90 && score <= 100 {
		return "A"
	} else if score >= 80 && score < 90 {
		return "B"
	} else if score >= 70 && score < 80 {
		return "C"
	} else if score >= 60 && score < 70 {
		return "D"
	} else if score < 60 && score > 0 {
		return "F"
	} else {
		return "Invalid"
	}
}
