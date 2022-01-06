package main

import "fmt"

func main() {
	// nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	// fmt.Printf("Array Sign: %d\n", arraySign(nums))

	// Merge Sort
	// nums := []int{5, 3, 4, 2, 1, 6}
	// nums = mergeSort(nums)
	// display(nums)

	// Search Rotated Array
	// nums := []int{3, 5, 1}
	// i := searchRotatedSorted(nums, 3)
	// fmt.Printf("Index: %d\n", i)

	// // Search Rotated Array Two (non-distinct values)
	// nums := []int{1, 3}
	// ret := searchRotatedSortedTwo(nums, 3)
	// fmt.Printf("Found: %t\n", ret)

	// // Binary Search
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// ret := binarySearch(nums, 3)
	// fmt.Printf("Found: %t\n", ret)

	// Search Insert
	// nums := []int{1, 3, 5, 6}
	// ret := searchInsert(nums, 6)
	// fmt.Printf("Index: %d\n", ret)

	// Number of Ways
	// nums := []int{1, 5, 3, 3, 3}
	// ret := numberOfWays(nums, 6)
	// fmt.Println(ret)

	// Are They Equal
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 4, 1, 1}
	// fmt.Println(areTheyEqual(a, b))

	n := make([][]int, 4)
	n[0] = []int{1, 5, 9}
	n[1] = []int{2, 6, 0}
	n[2] = []int{3, 7, 1}
	n[3] = []int{4, 8, 2}

	m := rotateMatrix90Degrees(n)
	fmt.Println(m)
}
