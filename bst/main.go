package main

import "fmt"

func main() {
	b := new(BST)
	b.insertRecursive(b.Root, 20)
	b.insertRecursive(b.Root, 10)
	b.insertRecursive(b.Root, 30)
	b.insertRecursive(b.Root, 25)
	b.insertRecursive(b.Root, 15)
	b.insertRecursive(b.Root, 5)
	b.insertRecursive(b.Root, 3)
	b.insertRecursive(b.Root, 11)

	inOrderRecursive(b.Root)
	fmt.Println()
	preOrderRecursive(b.Root)
	fmt.Println()
	postOrderRecursive(b.Root)
	fmt.Println()
}
