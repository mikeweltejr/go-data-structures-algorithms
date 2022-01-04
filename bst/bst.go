package main

import "fmt"

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
}

func (b *BST) insertRecursive(tempRoot *BSTNode, e int) *BSTNode {
	if tempRoot != nil {
		if e < tempRoot.Val {
			tempRoot.Left = b.insertRecursive(tempRoot.Left, e)
		} else {
			tempRoot.Right = b.insertRecursive(tempRoot.Right, e)
		}
	} else {
		newNode := new(BSTNode)
		newNode.Val = e
		tempRoot = newNode

		if b.Root == nil {
			b.Root = newNode
		}
	}

	return tempRoot
}

func inOrderRecursive(tempRoot *BSTNode) {
	if tempRoot != nil {
		inOrderRecursive(tempRoot.Left)
		fmt.Printf("%d->", tempRoot.Val)
		inOrderRecursive(tempRoot.Right)
	}
}

func preOrderRecursive(tempRoot *BSTNode) {
	if tempRoot != nil {
		fmt.Printf("%d->", tempRoot.Val)
		preOrderRecursive(tempRoot.Left)
		preOrderRecursive(tempRoot.Right)
	}
}

func postOrderRecursive(tempRoot *BSTNode) {
	if tempRoot != nil {
		postOrderRecursive(tempRoot.Left)
		postOrderRecursive(tempRoot.Right)
		fmt.Printf("%d->", tempRoot.Val)
	}
}

func search(tempRoot *BSTNode, n int) bool {
	if tempRoot == nil {
		return false
	}

	if tempRoot.Val == n {
		return true
	} else if n < tempRoot.Val {
		return search(tempRoot.Left, n)
	} else if n > tempRoot.Val {
		return search(tempRoot.Right, n)
	}

	return false
}
