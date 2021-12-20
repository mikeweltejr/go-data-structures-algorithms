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

func getUniformIntegerCountInInterval(a int, b int) int {
	count := 0

	if a == b {
		return 1
	}

	if a < 10 {
		count = 10 - a
	}

	return count
}
