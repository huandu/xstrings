// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"unicode"
	"unicode/utf8"
)

// Get str's utf8 rune length.
func Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Count number of words in a string.
//
// Word is defined as a locale dependent string containing alphabetic characters,
// which may also contain but not start with `'` and `-` characters.
func WordCount(str string) int {
	var r rune
	var size, n int

	inWord := false

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		switch {
		case isAlphabet(r):
			if !inWord {
				inWord = true
				n++
			}

		case inWord && (r == '\'' || r == '-'):
			// Still in word.

		default:
			inWord = false
		}

		str = str[size:]
	}

	return n
}

const minCJKCharacter = '\u3400'

// Checks r is a letter but not CJK character.
func isAlphabet(r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}

	switch {
	// Quick check for non-CJK character.
	case r < minCJKCharacter:
		return true

	// Common CJK characters.
	case r >= '\u4E00' && r <= '\u9FCC':
		return false

	// Rare CJK characters.
	case r >= '\u3400' && r <= '\u4D85':
		return false

	// Rare and historic CJK characters.
	case r >= '\U00020000' && r <= '\U0002B81D':
		return false
	}

	return true
}
