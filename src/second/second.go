package second

import (
	"log"
)

func isIsomorphic(s string, t string) bool {
	sMatchMap := make(map[byte]byte)
	tMatchMap := make(map[byte]byte)

	for i := 0; i < len(s); i += 1 {
		if val, ok := sMatchMap[s[i]]; ok {
			if val != t[i] {
				return false
			}
		}

		sMatchMap[s[i]] = t[i]
	}

	for i := 0; i < len(t); i += 1 {
		if val, ok := tMatchMap[t[i]]; ok {
			if val != s[i] {
				return false
			}
		}

		tMatchMap[t[i]] = s[i]
	}

	return true
}

func isSubsequence(s string, t string) bool {
	i := 0

	// Go along both of the string
	// If the character inside the longer string appears in the shorter string
	// Go to the next character of shorter string
	// Else moving on to search for in longer string
	for j := 0; j < len(t) && i < len(s); j += 1 {
		if t[j] == s[i] {
			i += 1
		}
	}

	return i == len(s)
}

func Run() {
	s := "abc"
	t := "ahbgdc"

	// result := isIsomorphic(s, t)
	result := isSubsequence(s, t)
	log.Println(result)
}
