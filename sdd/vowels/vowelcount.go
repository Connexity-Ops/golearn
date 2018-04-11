package vowels

import (
	"fmt"
	"strings"
)

//VowelCount given a string, adds together unique vowel occurances and returns a map of sums
func VowelCount(s string) map[string]int {

	c := 0 //initialize number of consonants
	y := 0 //initialize number of 'y's
	chars := s

	//initialize map to store vowel counts
	vowelMap := map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}

	//iterate over string argument and add up each occurance of a vowel
	for _, char := range strings.ToLower(chars) {
		switch char {
		case 'a':
			vowelMap["a"]++
		case 'e':
			vowelMap["e"]++
		case 'i':
			vowelMap["i"]++
		case 'o':
			vowelMap["o"]++
		case 'u':
			vowelMap["u"]++
		case 'y': //because sometimes "y" :)
			y++
			//vowelMap["y"]++
		default:
			c++
		}

	}

	//format and print output to show work
	fmt.Printf("\nString: %v\nConsonants: %v\nYs: %v\n", s, c, y)
	fmt.Print(vowelMap)
	fmt.Printf("\n")

	return vowelMap
}