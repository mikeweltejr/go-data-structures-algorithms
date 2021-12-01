package main

import (
	"math"
)

func isPrime(n int) bool {
	for x := 2; x < int(math.Sqrt(float64(n))); x++ {
		if n%x == 0 {
			return false
		}
	}

	return true
}
