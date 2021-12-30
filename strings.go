package main

import (
	"fmt"
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
