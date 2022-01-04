package main

import "fmt"

func main() {
	rotationalCipher("Zebra-493", 3)

	s := "abcdefghijklmnopqrstuvwxy211"
	isUnique := hasUniqueCharacters(s)
	fmt.Printf("Is %v Unique: %v\n", s, isUnique)

	s = urlify("Mr John Smith      ")
	fmt.Printf("URLIfied String: %v\n", s)

	s = "Tact coa"
	retBool := permutationHasPalindrome(s)
	fmt.Printf("Is Permutation Palindrome of %v: %v\n", s, retBool)

	a := "pale"
	b := "bake"
	r := oneAway(a, b)
	fmt.Printf("One Awway %v, %v: %v\n", a, b, r)
}
