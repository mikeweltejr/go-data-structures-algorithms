package main

import "fmt"

func main() {
	l := LinkedList{}

	l.addLast(1)
	l.addLast(2)
	l.addLast(3)
	l.addLast(4)
	l.addLast(5)
	l.addFirst(10)
	l.addAny(12, 4)

	l.removeFirst()
	l.removeLast()
	l.removeAny(0)

	l.display()

	fmt.Printf("Element found at position: %d\n", l.search(4))
}
