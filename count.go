// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"unicode/utf8"
)

// Get str's utf8 rune length.
func Len(str string) int {
	return utf8.RuneCountInString(str)
}
