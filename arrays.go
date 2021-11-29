package main

import "fmt"

func display(nums []int) {
	for _, n := range nums {
		fmt.Printf("%d,", n)
	}
	fmt.Println()
}

func arraySign(nums []int) int {
	c := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}

		if nums[i] < 0 {
			c++
		}
	}

	if c%2 == 1 {
		return -1
	}

	return 1
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// merge left
	left := mergeSort(nums[:len(nums)/2])
	// merge right
	right := mergeSort(nums[len(nums)/2:])

	return merge(left, right)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

/*

Solution:

	* Find the start of the sorted array
	* The end is start - 1 (if start is 0 it's already sorted)
	* Once we know start and end we can say if target < nums[start] return -1 since it won't be found
	* Similarly for end if target is greater than nums[end] return -1
	* Find out which the target is closer to start or end and search from there since it will be quicker
	* It's possible that the start can wrap (e.g. array with length 4 start is index 3 but it wraps to indeces 0,1)
		- in this case we reset the loop to 0 when it hits the array length so that we can continue searching
		- to prevent infinite we break when it hits the end as it is not found
	* We do the same thing for the end where it could potentially wrap (e.g. end is 0)
		- we reset the loop to len(nums) when it hits 0 to continue searching
		- to prevent infinite loop we stop when we hit the start
*/
func searchRotatedSorted(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	start := 0

	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if j > 0 && nums[j] < nums[j-1] {
			start = j
			break
		} else if !(i+1 >= len(nums)) && nums[i] > nums[i+1] {
			start = i + 1
			break
		}
	}

	end := len(nums) - 1
	if start != 0 {
		end = start - 1
	}

	fmt.Printf("Start: %d\n", start)
	fmt.Printf("End: %d\n", end)

	if target < nums[start] {
		return -1
	}

	if target > nums[end] {
		return -1
	}

	if target-nums[start] <= nums[end]-target {
		for i := start; i < len(nums); i++ {
			fmt.Println(i)
			if i == end {
				break
			}

			if nums[i] == target {
				return i
			}
			if i == len(nums)-1 {
				i = -1
			}
		}
	} else {
		for j := end; j >= 0; j-- {
			if j == start {
				break
			}

			if nums[j] == target {
				return j
			}

			if j == 0 {
				j = len(nums)
			}
		}
	}

	return -1
}

/*
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.
*/

/*

Solution:
	* Find start and end (this is more complicated as duplicates can exist)
	* To find the start we should keep going until nums[start] != nums[start-1]
	* If we can find the first occurrence of start we can get end very easily
	* Once we get these everything else should be just as the method above

*/
func searchRotatedSortedTwo(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return target == nums[0]
	}

	start := 0

	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if j > 0 && nums[j] < nums[j-1] {
			start = j
			break
		} else if !(i+1 >= len(nums)) && nums[i] > nums[i+1] {
			start = i + 1
			break
		}
	}

	minVal := nums[start]

	end := 0

	if start != 0 {
		end = start - 1
	} else {
		end = len(nums) - 1
	}
	// find the maxVal which will be the end of the sorted array
	for nums[end] == minVal && end != start {
		end--
	}

	fmt.Printf("Start: %d\n", start)
	fmt.Printf("End: %d\n", end)

	// if the target is less than the value of start or greater than value at end return false
	if target < nums[start] || target > nums[end] {
		return false
	}

	// if start and end are equal just return if the start is equal to target since the array is all the same
	if nums[start] == nums[end] {
		return nums[start] == target
	}

	if target-nums[start] <= nums[end]-target {
		for i := start; i < len(nums); i++ {
			fmt.Println(i)
			if i == end {
				break
			}

			if nums[i] == target {
				return true
			}
			if i == len(nums)-1 {
				i = -1
			}
		}
	} else {
		for j := end; j >= 0; j-- {
			if j == start {
				break
			}

			if nums[j] == target {
				return true
			}

			if j == 0 {
				j = len(nums)
			}
		}
	}

	return false
}
