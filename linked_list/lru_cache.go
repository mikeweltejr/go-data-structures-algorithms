package main

import "fmt"

type LRUCache struct {
	capacity int
	list     *DoubleLinkedList
	elements map[interface{}]*DblNode
}

func newLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     &DoubleLinkedList{},
		elements: make(map[interface{}]*DblNode, capacity),
	}
}

func (cache *LRUCache) Get(x interface{}) interface{} {
	node, ok := cache.elements[x]

	if ok {
		cache.list.MoveNodeToFront(node)
		return node.Element
	}

	return -1

}

func (cache *LRUCache) Put(x interface{}) {
	// If element does not exist add it to the map and LinkedList
	node, ok := cache.elements[x]
	if !ok {
		newNode := &DblNode{
			Element: x,
		}
		cache.elements[x] = newNode
		cache.list.AddFirst(newNode)
	} else {
		cache.list.MoveNodeToFront(node)
	}

	// If cache is full remove last recently used
	if cache.list.GetSize() > cache.capacity {
		cache.list.Tail = cache.list.Tail.Prev
		cache.list.Tail.Next = nil
	}

}

func (cache *LRUCache) PrintCache() {
	p := cache.list.Head

	for p != nil {
		fmt.Printf("%v-->", fmt.Sprint(p.Element))
		p = p.Next
	}
	fmt.Println()
}
