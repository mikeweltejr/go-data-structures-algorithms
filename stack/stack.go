package main

type Stack []int

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(x int) {
	*s = append(*s, x)
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
