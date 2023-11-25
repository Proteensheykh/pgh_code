package main

import (
	"fmt"
)

func main() {
	var char byte = 'd'

	if isVowel(char) {
		fmt.Println("Letter", string(char), "is a vowel")
	} else {
		fmt.Println("Letter", string(char), "is not a vowel")
	}

}

func isVowel(char byte) bool {
	switch char {
	case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
		return true
	default:
		return false
	}
}
