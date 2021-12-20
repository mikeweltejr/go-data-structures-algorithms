package main

type Queue []int

func (q *Queue) Enqueue(n int) {
	if len(*q) == 0 {
		*q = []int{}
		*q = append(*q, n)
		return
	}

	*q = append(*q, n)
}

func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		panic("Queue is empty")
	}

	n := (*q)[0]
	*q = (*q)[1:]

	return n
}
