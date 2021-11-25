package main

import "fmt"

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
		fmt.Printf("%d,", p.Element)

		p = p.Next
	}

	fmt.Println()
}