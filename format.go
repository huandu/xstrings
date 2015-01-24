// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"bytes"
	"unicode/utf8"
)

// ExpandTabs can expand tabs ('\t') rune in str to one or more spaces dpending on
// current column and tabSize.
// The column number is reset to zero after each newline ('\n') occurring in the str.
//
// ExpandTabs uses RuneWidth to decide rune's width.
// For example, CJK characters will be treated as two characters.
//
// If tabSize <= 0, ExpandTabs panics with error.
//
// Samples:
//     ExpandTabs("a\tbc\tdef\tghij\tk", 4) => "a   bc  def ghij    k"
//     ExpandTabs("abcdefg\thij\nk\tl", 4)  => "abcdefg hij\nk   l"
//     ExpandTabs("z中\t文\tw", 4)           => "z中 文  w"
func ExpandTabs(str string, tabSize int) string {
	if tabSize <= 0 {
		panic("tab size must be positive")
	}

	var r rune
	var i, size, column, expand int
	var output *bytes.Buffer

	orig := str

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		if r == '\t' {
			expand = tabSize - column%tabSize

			if output == nil {
				output = allocBuffer(orig, str)
			}

			for i = 0; i < expand; i++ {
				output.WriteByte(byte(' '))
			}

			column += expand
		} else {
			if r == '\n' {
				column = 0
			} else {
				column += RuneWidth(r)
			}

			if output != nil {
				output.WriteRune(r)
			}
		}

		str = str[size:]
	}

	if output == nil {
		return orig
	}

	return output.String()
}
