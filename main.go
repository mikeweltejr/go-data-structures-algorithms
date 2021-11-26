package main

import "fmt"

func main() {
	l := LinkedList{}

	l.insertSorted(12)
	l.insertSorted(1)
	l.insertSorted(3)
	l.insertSorted(5)
	l.insertSorted(12)
	l.insertSorted(11)
	l.insertSorted(20)
	l.insertSorted(3)
	l.insertSorted(15)
	l.display()

	l.reverseList()
	l.display()
	//l.move(1)

	// l.removeFirst()
	// l.removeLast()
	// l.removeAny(0)

	fmt.Printf("Element found at position: %d\n", l.search(4))
	fmt.Printf("Has Cycle? %t\n", l.hasCycle())
}
