// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"strconv"
	"strings"
	"testing"
)

func TestExpandTabs(t *testing.T) {
	runner := func(str string) (result string) {
		defer func() {
			if e := recover(); e != nil {
				result = e.(string)
			}
		}()

		input := strings.Split(str, separator)
		n, _ := strconv.Atoi(input[1])
		return ExpandTabs(input[0], n)
	}

	runTestCases(t, runner, _M{
		sep("a\tbc\tdef\tghij\tk", "4"): "a   bc  def ghij    k",
		sep("abcdefg\thij\nk\tl", "4"):  "abcdefg hij\nk   l",
		sep("z中\t文\tw", "4"):            "z中 文  w",
		sep("abcdef", "4"):              "abcdef",

		sep("abc\td\tef\tghij\nk\tl", "3"): "abc   d  ef ghij\nk  l",
		sep("abc\td\tef\tghij\nk\tl", "1"): "abc d ef ghij\nk l",

		sep("abc", "0"):  "tab size must be positive",
		sep("abc", "-1"): "tab size must be positive",
	})
}
