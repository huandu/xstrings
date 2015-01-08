// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"strings"
	"unicode/utf8"
)

// Reverse a utf8 encoded string.
func Reverse(str string) string {
	var size int

	tail := len(str)
	buf := make([]byte, tail)
	s := buf

	for len(str) > 0 {
		_, size = utf8.DecodeRuneInString(str)
		tail -= size
		s = append(s[:tail], []byte(str[:size])...)
		str = str[size:]
	}

	return string(buf)
}

// Slice a string by rune.
// Start and end must satisfy 0 <= start <= end <= rune length.
// Otherwise, Slice will panic as out of range.
func Slice(str string, start, end int) string {
	var size, startPos, endPos int

	origin := str

	if start < 0 || end > len(str) || start > end {
		panic("out of range")
	}

	end -= start

	for start > 0 && len(str) > 0 {
		_, size = utf8.DecodeRuneInString(str)
		start--
		startPos += size
		str = str[size:]
	}

	endPos = startPos

	for end > 0 && len(str) > 0 {
		_, size = utf8.DecodeRuneInString(str)
		end--
		endPos += size
		str = str[size:]
	}

	if len(str) == 0 && (start > 0 || end > 0) {
		panic("out of range")
	}

	return origin[startPos:endPos]
}

// Partition splits a string by sep into three parts.
// The return value is a slice of strings with head, match and tail.
//
// If str contains sep, for example "hello" and "l", Partition returns
//     []string{"he", "l", "lo"}
//
// If str doesn't contain sep, for example "hello" and "x", Partition returns
//     []string{"hello", "", ""}
func Partition(str, sep string) []string {
	index := strings.Index(str, sep)

	if index == -1 {
		return []string{str, "", ""}
	}

	return []string{str[:index], sep, str[index+len(sep):]}
}
