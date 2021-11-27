package main

func main() {
	l := LinkedList{}

	l.addOddEven(2)
	l.addOddEven(4)
	l.addOddEven(3)
	l.addOddEven(1)
	l.addOddEven(7)
	l.addOddEven(6)
	l.display()

	l.oddEventList()
	l.display()
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
