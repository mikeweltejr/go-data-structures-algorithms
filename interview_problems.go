package main

import "fmt"

func getMaximumEatenDishCount(N int32, D []int32, K int32) int32 {
	uniqueDishes := []int32{D[0]}

	for i := int32(1); i < N; i++ {
		isUnique := true

		for j := int32(1); j <= K; j++ {
			if i-j < 0 {
				break
			}

			if D[i] == uniqueDishes[len(uniqueDishes)-int(j)] {
				isUnique = false
				break
			}
		}

		if isUnique {
			uniqueDishes = append(uniqueDishes, D[i])
		}
	}

	fmt.Println(uniqueDishes)

	return int32(len(uniqueDishes))
}
