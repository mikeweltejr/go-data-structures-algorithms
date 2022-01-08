package main

import (
	"fmt"
	"strconv"
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

func (n *Node) display() {
	for n != nil {
		fmt.Printf("%d-->", n.Element)
		n = n.Next
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

func (l *LinkedList) reverseEvens() {
	if l.isEmpty() {
		return
	}

	current := l.Head
	var prev *Node

	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	// 1, 2, 4, 6, 9, 8, 12

	//

	prev.display()
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

func (l *LinkedList) hasCycleTwo() bool {
	if l.isEmpty() {
		return false
	}

	nodeMap := make(map[*Node]bool)
	p := l.Head

	for p != nil {
		if nodeMap[p.Next] {
			return true
		}
		nodeMap[p] = true
		p = p.Next
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

func (l *LinkedList) findIntersection(l2 *LinkedList) *Node {
	nodeMap := make(map[*Node]*Node)

	p := l.Head

	for p != nil {
		nodeMap[p] = p
		p = p.Next
	}

	p2 := l2.Head

	for p2 != nil {
		if nodeMap[p2] == p2 {
			return p2
		}
		p2 = p2.Next
	}

	return nil
}

func (head *Node) removeDuplicates() *Node {
	nodeMap := make(map[int]bool)
	p := head
	var prev *Node

	for p != nil {
		if nodeMap[p.Element] {
			prev.Next = p.Next
		} else {
			prev = p
			nodeMap[p.Element] = true
		}

		p = p.Next
	}

	return head
}

func (head *Node) removeDuplicatesDistinct() *Node {
	var prev *Node
	p := head

	for p != nil {
		if p.Next == nil {
			if prev != nil {
				prev.Next = p
			} else {
				head = p
			}
			return head
		}

		if p.Element == p.Next.Element {
			for p.Next != nil && p.Element == p.Next.Element {
				p = p.Next
			}
			p = p.Next

			if p == nil {
				if prev != nil {
					prev.Next = nil
				} else {
					head = nil
				}
			}
		} else {
			if prev == nil {
				head = p
				prev = head
			} else {
				prev.Next = p
				prev = p
			}

			p = p.Next
		}
	}

	return head
}

func (head *Node) deleteDuplicatesUnsorted() *Node {
	dupMap := make(map[int]bool)
	p := head

	for p != nil {
		_, ok := dupMap[p.Element]
		if ok {
			dupMap[p.Element] = true
		} else {
			dupMap[p.Element] = false
		}
		p = p.Next
	}

	p = head
	var prev *Node

	for p != nil {
		if !dupMap[p.Element] {
			if prev == nil {
				head = p
				prev = head
			} else {
				prev.Next = p
				prev = p
			}
		} else if p.Next == nil {
			if prev != nil {
				prev.Next = nil
			}
		}
		p = p.Next
	}

	if prev == nil {
		return nil
	}

	return head
}

func swapPairs(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

func (head *Node) returnNthNode(n int) *Node {
	nodeMap := make(map[int]*Node)
	i := 0
	p := head

	for p != nil {
		nodeMap[i] = p
		p = p.Next
		i++
	}

	if len(nodeMap) < n {
		return nil
	}

	return nodeMap[len(nodeMap)-n]
}

func (head *Node) partition(n int) *Node {
	p := head
	var leftHead *Node
	var rightHead *Node
	var lastLeft *Node

	for p != nil {
		newNode := Node{
			Element: p.Element,
			Next:    nil,
		}
		if p.Element < n {
			if leftHead == nil {
				leftHead = &newNode
				lastLeft = leftHead
			} else {
				newNode.Next = leftHead
				leftHead = &newNode
			}
		} else {
			if rightHead == nil {
				rightHead = &newNode
			} else {
				newNode.Next = rightHead
				rightHead = &newNode
			}
		}
		p = p.Next
	}

	if leftHead == nil {
		leftHead = rightHead
	} else {
		lastLeft.Next = rightHead
	}

	return leftHead
}

func sumReversed(a *Node, b *Node) *Node {
	aSum := 0
	ind := 1

	p := a
	for p != nil {
		aSum += (p.Element * ind)
		ind *= 10
		p = p.Next
	}

	p = b
	bSum := 0
	ind = 1
	for p != nil {
		bSum += (p.Element * ind)
		ind *= 10
		p = p.Next
	}

	retSum := aSum + bSum
	s := strconv.Itoa(retSum)
	var retNode *Node

	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(string(s[i]))
		newNode := Node{
			Element: n,
			Next:    nil,
		}
		if retNode == nil {
			retNode = &newNode
		} else {
			newNode.Next = retNode
			retNode = &newNode
		}
	}

	return retNode
}

func sumForward(a *Node, b *Node) *Node {
	p := a
	aArr := []int{}
	for p != nil {
		aArr = append(aArr, p.Element)
		p = p.Next
	}

	bArr := []int{}
	p = b
	for p != nil {
		bArr = append(bArr, p.Element)
		p = p.Next
	}

	aSum := 0
	ind := 1
	for i := len(aArr) - 1; i >= 0; i-- {
		aSum += (aArr[i] * ind)
		ind *= 10
	}

	bSum := 0
	ind = 1
	for i := len(bArr) - 1; i >= 0; i-- {
		bSum += (bArr[i] * ind)
		ind *= 10
	}

	retSum := aSum + bSum
	var retNode *Node
	s := strconv.Itoa(retSum)

	for i := len(s) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(s[i]))
		newNode := Node{
			Element: n,
			Next:    nil,
		}
		if retNode == nil {
			retNode = &newNode
		} else {
			newNode.Next = retNode
			retNode = &newNode
		}
	}

	return retNode
}
