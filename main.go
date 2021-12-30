package main

import "fmt"

func main() {
	listExamples()
}

func stringExamples() {
	// rotationalCipher("Zebra-493", 3)

	// s := "abcdefghijklmnopqrstuvwxy211"
	// isUnique := hasUniqueCharacters(s)
	// fmt.Printf("Is %v Unique: %v\n", s, isUnique)
}

func listExamples() {
	intList := List[int]{}
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(4)
	intList.Add(5)
	fmt.Println(intList)
	intList.Remove(1)
	fmt.Println(intList)
}

// func LRUCacheExamples() {
// 	cache := newLRUCache(7)
// 	cache.Put(1)
// 	cache.Put(2)
// 	cache.Put(3)
// 	cache.Put(4)
// 	cache.Put(5)
// 	cache.Put(6)
// 	cache.PrintCache()

// 	cache.Get(4)
// 	cache.PrintCache()

// 	cache.Get(3)
// 	cache.PrintCache()

// 	cache.Put(7)
// 	cache.PrintCache()

// 	cache.Put(8)
// 	cache.PrintCache()

// 	cache.Get(9)
// 	cache.PrintCache()

// 	cache.Get(7)
// 	cache.PrintCache()
// }

// func queueExamples() {
// 	var q Queue
// 	q.Enqueue(1)
// 	q.Enqueue(2)
// 	q.Enqueue(3)
// 	q.Enqueue(4)
// 	q.Enqueue(5)

// 	fmt.Println(q.Dequeue())
// 	fmt.Println(q.Dequeue())
// 	fmt.Println(q.Dequeue())
// 	fmt.Println(q.Dequeue())
// 	fmt.Println(q.Dequeue())
// }

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
}

func numbersExample() {
	// num := 31
	// ret := isPrime(num)
	// fmt.Printf("Is %d Prime? %t\n", num, ret)
}

func linkedListExamples() {
	//l := LinkedList{}

	// l.addOddEven(2)
	// l.addOddEven(4)
	// l.addOddEven(3)
	// l.addOddEven(1)
	// l.addOddEven(7)
	// l.addOddEven(6)
	//l.oddEventList()

	// l.addLast(1)
	// l.addLast(2)
	// l.addLast(3)
	// l.addLast(4)
	// l.addLast(5)
	// l.addLast(6)
	// l.display()

	// n := swapPairs(l.Head)
	// n.display()

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

	// l := LinkedList{}
	// l.addLast(1)
	// l.addLast(2)
	// l.addLast(4)
	// l.addLast(6)
	// l.addLast(9)
	// l.addLast(8)
	// l.addLast(12)
	// l.reverseEvens()

	// l := LinkedList{}
	// l.addFirst(2)
	// l.addFirst(3)
	// l.addFirst(1)
	// l.addFirst(4)
	// l.addFirst(5)
	// l.addFirst(7)
	// l.addFirst(2)
	// l.addFirst(3)
	// l.addFirst(5)
	// l.addFirst(6)
	// l.display()
	// n := l.Head.returnNthNode(5)
	// n.display()
}

func graphExamples() {
	// g := new(Graph)
	// g.initialize(true)

	// g.addEdge(0, 1)
	// g.addEdge(0, 2)
	// g.addEdge(1, 2)
	// g.addEdge(2, 0)
	// g.addEdge(2, 3)
	// g.addEdge(3, 3)

	// g.dfs(2)
}

func interviewProblems() {
	// n := getMaximumEatenDishCount(6, []int32{1, 2, 3, 3, 2, 1}, 6)
	// fmt.Println(n)
}
