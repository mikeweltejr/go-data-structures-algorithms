package main

type List[T comparable] []T

func (l *List[T]) Count() int {
	return len(*l)
}

func(l *List[T]) Add(item T) {
	if l.Count() == 0 {
		*l = []T{} 
	}

	*l = append(*l, item)
}

func (l *List[T]) Remove(item T) {
	if l.Count() == 0 {
		panic("List is empty")
	}

	// Need to figure out how to implement an Equals for different types
	j := 0
	newArr := make([]T, l.Count())
	for _, i := range *l {
		if i != item {
			newArr[j] = i
			j++
		}
	}

	*l = newArr[:j]
}
