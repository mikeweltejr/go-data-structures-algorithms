package main

func main() {
	linkedListExamples()
}

func arrayExamples() {
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
}

func numbersExample() {
	// num := 31
	// ret := isPrime(num)
	// fmt.Printf("Is %d Prime? %t\n", num, ret)
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

	n := swapPairs(l.Head)
	n.display()

	// fmt.Printf("Middle Node: %d\n", l.findMiddleNode().Element)
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
