// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"strconv"
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	runTestCases(t, Reverse, map[string]string{
		"reverse string": "gnirts esrever",
		"中文如何？":          "？何如文中",
		"中en文混~排怎样？a":    "a？样怎排~混文ne中",
	})
}

func TestSlice(t *testing.T) {
	SliceRunner := func(str string) (result string) {
		defer func() {
			if e := recover(); e != nil {
				result = e.(string)
			}
		}()

		strs := strings.Split(str, " ¶ ")
		start, _ := strconv.ParseInt(strs[1], 10, 0)
		end, _ := strconv.ParseInt(strs[2], 10, 0)

		result = Slice(strs[0], int(start), int(end))
		return
	}

	runTestCases(t, SliceRunner, map[string]string{
		"abcdefghijk ¶ 3 ¶ 8":      "defgh",
		"来点中文如何？ ¶ 2 ¶ 7":          "中文如何？",
		"中en文混~排总是少不了的a ¶ 2 ¶ 8":   "n文混~排总",
		"中en文混~排总是少不了的a ¶ 0 ¶ 0":   "",
		"中en文混~排总是少不了的a ¶ 14 ¶ 14": "",

		"let us slice out of range ¶ -3 ¶ 3": "out of range",
		"超出范围哦 ¶ 2 ¶ 6":                      "out of range",
		"don't do this ¶ 3 ¶ 2":              "out of range",
		"千gan万de不piao要liang ¶ 19 ¶ 19":       "out of range",
	})
}
