// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	runner := func(str string) string {
		return fmt.Sprint(Len(str))
	}

	runTestCases(t, runner, _M{
		"abcdef":    "6",
		"中文":        "2",
		"中yin文hun排": "9",
		"":          "0",
	})
}

func TestWordCount(t *testing.T) {
	runner := func(str string) string {
		return fmt.Sprint(WordCount(str))
	}

	runTestCases(t, runner, _M{
		"one word: λ":             "3",
		"中文":                      "0",
		"你好，sekai！":               "1",
		"oh, it's super-fancy!!a": "4",
		"":        "0",
		"-":       "0",
		"it's-'s": "1",
	})
}
