package main

import (
	"fmt"
	"sort"
)

type Stack []int

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(x int) {
	*s = append(*s, x)
}

func (s *Stack) pushMin(x int) {
	if s.isEmpty() {
		*s = append(*s, x)
		return
	}

	index := len(*s) - 1
	element := (*s)[index]

	if x > element {
		(*s)[index] = x
		*s = append(*s, element)
	} else {
		*s = append(*s, x)
	}
}

func (s *Stack) pop() (int, bool) {
	if s.isEmpty() {
		return -1, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}

func (s *Stack) min() int {
	return (*s)[len(*s)-1]
}

// TODO: can we do this without an int array, and only another stack?
func (s *Stack) sort() *Stack {
	n := []int{}

	for !s.isEmpty() {
		x, _ := s.pop()
		n = append(n, x)
	}

	sort.Ints(n)

	for i := len(n) - 1; i >= 0; i-- {
		s.push(n[i])
	}

	return s
}

func (s *Stack) display() {
	for !s.isEmpty() {
		n, _ := s.pop()
		fmt.Printf("%d, ", n)
	}
	fmt.Println()
}
