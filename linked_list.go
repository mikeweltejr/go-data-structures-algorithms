package main

import (
	"fmt"
)

type Node struct {
	Element int
	Next    *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) getSize() int {
	return l.Size
}

func (l *LinkedList) isEmpty() bool {
	return l.getSize() == 0
}

func (l *LinkedList) addLast(e int) {
	newNode := Node{
		Element: e,
		Next:    nil,
	}

	if l.isEmpty() {
		l.Head = &newNode
	} else {
		l.Tail.Next = &newNode
	}

	l.Tail = &newNode
	l.Size = l.getSize() + 1
}

func (l *LinkedList) addFirst(e int) {
	newNode := Node{
		Element: e,
		Next:    nil,
	}

	if l.isEmpty() {
		l.Head = &newNode
		l.Tail = &newNode
	} else {
		newNode.Next = l.Head
		l.Head = &newNode
	}

	l.Size = l.getSize() + 1
}

func (l *LinkedList) addAny(e int, position int) {
	if position <= 0 || position > l.getSize() {
		fmt.Printf("Invalid Position: %d\n", position)
	}

	newNode := Node{
		Element: e,
		Next:    nil,
	}

	p := l.Head
	i := 1

	for i < position-1 {
		p = p.Next
		i++
	}

	newNode.Next = p.Next
	p.Next = &newNode

	l.Size = l.getSize() + 1
}

func (l *LinkedList) removeFirst() int {
	if l.isEmpty() {
		fmt.Println("Linked List is empty")
		return -1
	}

	element := l.Head.Element

	l.Head = l.Head.Next
	l.Size = l.getSize() - 1

	return element
}

func (l *LinkedList) removeLast() int {
	if l.isEmpty() {
		fmt.Println("Linked List is empty")
		return -1
	}

	element := l.Tail.Element

	p := l.Head
	i := 1

	for i < l.getSize()-1 {
		p = p.Next
		i++
	}

	p.Next = nil
	l.Tail = p
	l.Size = l.getSize() - 1

	return element
}

func (l *LinkedList) removeAny(position int) int {
	if position > l.getSize() || position < 1 {
		fmt.Printf("Invalid Position: %d\n", position)
		return -1
	}

	p := l.Head
	i := 1

	for i < position-1 {
		p = p.Next
		i++
	}

	e := p.Next.Element
	p.Next = p.Next.Next
	l.Size = l.getSize() - 1

	return e
}

func (l *LinkedList) search(e int) int {
	p := l.Head
	i := 0

	for p != nil {
		if p.Element == e {
			return i
		}

		p = p.Next
		i++
	}

	return -1
}

func (l *LinkedList) display() {
	p := l.Head

	for p != nil {
		fmt.Printf("%d ------>", p.Element)

		p = p.Next
	}

	fmt.Println()
}

func (l *LinkedList) addOddEven(e int) {
	if e%2 == 0 {
		l.addLast(e)
	} else {
		l.addFirst(e)
	}
}

func (l *LinkedList) insertSorted(e int) {
	newNode := Node{
		Element: e,
		Next:    nil,
	}

	if l.isEmpty() {
		l.Head = &newNode
		l.Tail = &newNode
	} else {
		p := l.Head
		q := l.Head

		for p != nil && p.Element < e {
			q = p
			p = p.Next
		}

		if p == l.Head {
			newNode.Next = l.Head
			l.Head = &newNode
		} else {
			newNode.Next = q.Next
			q.Next = &newNode
		}
	}

	temp := l.Head

	for temp.Next != nil {
		temp = temp.Next
	}

	l.Tail = temp

	l.Size = l.getSize() + 1
}

// TODO: fix this method
func (l *LinkedList) move(k int) {

	// If empty or k is 0 nothing to move so return
	if l.isEmpty() || k == 0 {
		return
	}

	temp := l.Tail

	// Set k to be something that can be properly moved given the size of the list (e.g. 6%5=1)
	if k > l.getSize() {
		k = k % l.getSize()
	}

	current := l.Head
	count := 1

	for count < k && current != nil {
		current = current.Next
		count++
	}

	if current == nil {
		return
	}

	fmt.Println(current.Element)

	kthNode := current
	temp.Next = l.Head
	l.Head = kthNode.Next
	kthNode.Next = nil
}

func (l *LinkedList) reverseList() {
	if l.isEmpty() {
		return
	}

	var prev *Node
	current := l.Head

	for current != nil {
		temp := current.Next
		// Takes each node and sets it's Next to the previous Node
		current.Next = prev
		prev = current
		current = temp
	}

	l.Head = prev
}

func (l *LinkedList) hasCycle() bool {
	if l.isEmpty() {
		return false
	}

	slow := l.Head
	fast := l.Head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func (l *LinkedList) isPalindrome() bool {
	p := l.Head
	intArr := []int{}

	for p != nil {
		intArr = append(intArr, p.Element)
		p = p.Next
	}

	if len(intArr) == 1 {
		return true
	}

	for i, j := 0, len(intArr)-1; i < len(intArr); i, j = i+1, j-1 {
		if intArr[i] != intArr[j] {
			return false
		}
	}

	return true
}

func (l *LinkedList) removeElements(val int) {
	if l.isEmpty() {
		return
	}

	p := l.Head
	var prev *Node

	for p != nil {
		if p.Element == val {
			if prev != nil {
				prev.Next = p.Next
				p = prev
			}
		}

		prev = p
		p = p.Next
	}

	if l.Head.Element == val {
		l.Head = l.Head.Next
	}
}

func (l *LinkedList) oddEventList() {
	if l.Head == nil {
		return
	}

	odd := l.Head
	even := l.Head.Next
	evenHead := even

	// 7 --> 1 --> 3 --> 2 --> 4 --> 6
	for even != nil && even.Next != nil {
		// 3, 4
		odd.Next = even.Next
		// 3, 4
		odd = odd.Next
		// 2, 6
		even.Next = odd.Next
		// 2, 6
		even = even.Next
	}

	odd.Next = evenHead
}

func (l *LinkedList) findMiddleNode() *Node {
	if l.isEmpty() {
		return nil
	}
	nodeArr := []*Node{}

	p := l.Head

	for p != nil {
		nodeArr = append(nodeArr, p)
		p = p.Next
	}

	mid := len(nodeArr) / 2

	return nodeArr[mid]
}
