package main

import "fmt"

func main() {
	arrayExamples()
}

func arrayExamples() {
	// nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	// fmt.Printf("Array Sign: %d\n", arraySign(nums))

	// Merge Sort
	// nums := []int{5, 3, 4, 2, 1, 6}
	// nums = mergeSort(nums)
	// display(nums)

	// Search Rotated Array
	nums := []int{3, 5, 1}
	i := searchRotatedSorted(nums, 3)
	fmt.Printf("Index: %d\n", i)
}

func linkedListExamples() {
	l := LinkedList{}

	// l.addOddEven(2)
	// l.addOddEven(4)
	// l.addOddEven(3)
	// l.addOddEven(1)
	// l.addOddEven(7)
	// l.addOddEven(6)
	//l.oddEventList()

	l.addLast(1)
	l.addLast(2)
	l.addLast(3)
	l.addLast(4)
	l.addLast(5)
	l.addLast(6)
	l.display()
	fmt.Printf("Middle Node: %d\n", l.findMiddleNode().Element)
	//l.reverseList()
	//l.display()
	//l.move(1)

	// l.removeFirst()
	// l.removeLast()
	// l.removeAny(0)

	// l.removeElements(7)
	// l.display()

	// fmt.Printf("Element found at position: %d\n", l.search(4))
	// fmt.Printf("Has Cycle? %t\n", l.hasCycle())
	// fmt.Printf("Is Palindrome? %t\n", l.isPalindrome())
}
