package main

type DblNode struct {
	Element interface{}
	Next    *DblNode
	Prev    *DblNode
}

type DoubleLinkedList struct {
	Head *DblNode
	Tail *DblNode
	Size int
}

func (l *DoubleLinkedList) IsEmpty() bool {
	return l.Size == 0
}

func (l *DoubleLinkedList) GetSize() int {
	return l.Size
}

func (l *DoubleLinkedList) AddFirst(node *DblNode) {
	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}

	l.Size = l.Size + 1
}

func (l *DoubleLinkedList) AddLast(e int) {
	newNode := &DblNode{
		Element: e,
	}

	if l.IsEmpty() {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
	}

	l.Size = l.Size + 1
}

func (l *DoubleLinkedList) AddAny(e int, position int) {
	if position <= 0 || position > l.GetSize() {
		panic("Position does not fit in with linked list size")
	}

	newNode := &DblNode{
		Element: e,
	}

	p := l.Head
	i := 1

	for i < position-1 {
		p = (*p).Next
		i++
	}

	newNode.Next = p.Next
	p.Next.Prev = newNode
	p.Next = newNode
	newNode.Prev = p

	l.Size = l.Size + 1
}

func (l *DoubleLinkedList) MoveNodeToFront(node *DblNode) {
	if l.IsEmpty() {
		l.AddFirst(node)
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Next = l.Head
		node.Prev = nil
		l.Head.Prev = node
		l.Head = node
	}
}
