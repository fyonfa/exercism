package raindrops

import (
	"strconv"
)

// Convert returns Plin, Plang or Plong depending if is a factor of 3, 5, 7
func Convert(num int) string {
	words := ""
	if num%3 == 0 {
		words += "Pling"
	}
	if num%5 == 0 {
		words += "Plang"
	}
	if num%7 == 0 {
		words += "Plong"
	}
	if words != "" {
		return words
	}
	return strconv.Itoa(num)
}
