package main

import "fmt"

func main() {
	// l := LinkedList{}

	// l.addOddEven(2)
	// l.addOddEven(4)
	// l.addOddEven(3)
	// l.addOddEven(1)
	// l.addOddEven(7)
	// l.addOddEven(6)
	// l.oddEventList()

	// l = LinkedList{}
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
	// l.reverseList()
	// l.display()

	// l.removeFirst()
	// l.removeLast()
	// l.removeAny(0)

	// l.removeElements(7)
	// l.display()

	// fmt.Printf("Element found at position: %d\n", l.search(4))
	// fmt.Printf("Has Cycle? %t\n", l.hasCycle())
	// fmt.Printf("Is Palindrome? %t\n", l.isPalindrome())

	// l = LinkedList{}
	// l.addLast(1)
	// l.addLast(2)
	// l.addLast(4)
	// l.addLast(6)
	// l.addLast(9)
	// l.addLast(8)
	// l.addLast(12)
	// l.reverseEvens()

	// l = LinkedList{}
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
	// n = l.Head.returnNthNode(5)
	// n.display()

	// l := LinkedList{}
	// l.addFirst(1)
	// l.addFirst(10)
	// l.addFirst(8)
	// l.addFirst(3)
	// l.addFirst(2)
	// l.addFirst(5)
	// l.addFirst(5)
	// n := l.Head.partition(5)
	// n.display()

	a := LinkedList{}
	a.addFirst(6)
	a.addFirst(1)
	a.addFirst(7)

	b := LinkedList{}
	b.addFirst(2)
	b.addFirst(9)
	b.addFirst(5)

	// n := sumReversed(a.Head, b.Head)
	// n.display()
	// n = sumForward(a.Head, b.Head)
	// n.display()

	b.Head.Next.Next = a.Head.Next
	fmt.Printf("Are a and b interseting: %v\n", intersecting(&a, &b))
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
