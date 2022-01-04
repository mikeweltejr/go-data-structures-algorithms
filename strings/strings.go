package main

import (
	"fmt"
	"math"
	"strings"
)

func rotationalCipher(input string, rotationalCipher int) string {
	retStr := ""

	for i := 0; i < len(input); i++ {
		c := input[i]

		if isNumber(int(c)) {
			c = c - '0'
			n := (int(c) + rotationalCipher) % 10
			retStr += fmt.Sprint(n)
		} else if isUpperAlpha(int(c)) {
			rotationalCipher = rotationalCipher % 26
			n := 65

			if int(c)+rotationalCipher > 90 {
				n = n + (int(c) + rotationalCipher) - 91
			} else {
				n = int(c) + rotationalCipher
			}
			char := rune(n)
			retStr += string(char)
		} else if isLowerAlpha(int(c)) {
			rotationalCipher = rotationalCipher % 26
			n := 97

			if int(c)+rotationalCipher > 122 {
				n = n + (int(c) + rotationalCipher) - 123
			} else {
				n = int(c) + rotationalCipher
			}
			char := rune(n)
			retStr += string(char)
		} else {
			retStr += string(c)
		}
	}

	fmt.Println(retStr)

	return retStr
}

func isNumber(n int) bool {
	return n >= 48 && n <= 57
}

func isUpperAlpha(n int) bool {
	return (n >= 65 && n <= 90)
}

func isLowerAlpha(n int) bool {
	return (n >= 97 && n <= 122)
}

func hasUniqueCharacters(s string) bool {
	charMap := make(map[rune]bool)
	s = strings.ToLower(s)

	for _, c := range s {
		if !charMap[rune(c)] {
			charMap[rune(c)] = true
		} else {
			return false
		}
	}

	return true
}

func urlify(s string) string {
	charArr := make([]rune, len(s))
	s = strings.TrimRight(s, " ")

	index := 0
	for i := 0; i < len(s); i++ {
		r := rune(s[i])

		if r == ' ' {
			charArr[index] = '%'
			charArr[index+1] = '2'
			charArr[index+2] = '0'
			index += 3
		} else {
			charArr[index] = r
			index++
		}
	}

	return string(charArr)
}

func permutationHasPalindrome(s string) bool {
	charMap := make(map[rune]int)
	charMatches := 0
	charCount := 0

	s = strings.ToLower(s)

	// Case when
	if len(s) <= 2 {
		return true
	}

	for i := 0; i < len(s); i++ {
		r := rune(s[i])

		if r != ' ' {
			charCount++
		}

		charMap[r] += 1

		if charMap[r] == 2 {
			charMatches += 1
			charMap[r] = 0
		}
	}

	if charCount/2 == charMatches {
		return true
	}

	return false
}

/*
	Returns true if one char can be inserted, removed, or updated to make the strings equal
*/
func oneAway(a string, b string) bool {
	if math.Abs(float64(len(a)))-math.Abs(float64(len(b))) > 1 {
		return false
	}

	maxString := a
	minString := b
	charDiff := 0

	if len(b) > len(a) {
		maxString = b
		minString = a
	}

	for i, j := 0, 0; j < len(maxString) && i < len(minString); {
		if minString[i] != maxString[j] {
			charDiff += 1

			if len(minString) == len(maxString) {
				i++
			}
		} else {
			i++
		}
		j++
	}

	return charDiff <= 1
}
