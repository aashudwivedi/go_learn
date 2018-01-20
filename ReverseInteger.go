package main

import (
	"fmt"
)

func reverse(x int) int {
	digits := make([] int, 1000)

	for x > 0 {
		digit := x % 10
		x = x / 10
		append(digits, digit)
	}
	return digits
}