package slices

import (
	"unicode"
)

func leftShift(s []byte, i, j int) []byte {
	n := j - i // number of spaces (with no first)
	for ; j < len(s); i++ {
		s[i] = s[j]
		j++
	}
	return s[:len(s)-n]
}

func SpaceMerge(s []byte) []byte {
	/*
		Merge sequence of Unicode spaces to one space of ASCII in byte slice
		(with no allocating additional memory)
	 */
	 s = s[:len(s)]
	for i, v1 := range s {
		if unicode.IsSpace(rune(v1)) {
			for j := i + 1; j < len(s); j++ {
				if !unicode.IsSpace(rune(s[j])) {
					s = leftShift(s, i+1, j)
					break
				}
				if j == len(s) - 1 {
					s = s[:i]
				}
			}
		}
	}
	if unicode.IsSpace(rune(s[len(s)-1])) {
		return s[:len(s)-1] // with no last space
	}
	return s
}
