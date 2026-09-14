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

		// Every invalid byte is counted as one rune.
		"\xff\xfe": "2",
		"a\xffb":   "3",

		// A rune and a following combining mark are two runes.
		"e\u0301": "2",

		"\uFFFD": "1",
		"😀a":     "2",
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
		"":                        "0",
		"-":                       "0",
		"it's-'s":                 "1",

		// A word may contain but not start with "'" or "-".
		"'":      "0",
		"''a''":  "1",
		"a-b":    "1",
		"a--b":   "1",
		"-a-":    "1",
		"don't":  "1",
		"you're": "1",

		// Numbers are not alphabetic, so they act as word separators.
		"123":       "0",
		"abc123def": "2",
		"a1b":       "2",

		// Underscore is not alphabetic either.
		"a_b": "2",

		// CJK ideographs are not words, but most other non-ASCII letters are.
		"中文abc":   "1",
		"こんにちは":   "1",
		"한국어":     "1",
		"ＡＢＣ":     "1",
		"Привет":  "1",
		"\uFFFD":  "0",
		"\xffabc": "1",

		// A combining mark is not a letter, so it breaks a word.
		"e\u0301x": "2",

		// Boundaries of the CJK ranges which are not treated as words.
		"\u3400":     "0",
		"\u4d85":     "0",
		"\u4e00":     "0",
		"\u9fcc":     "0",
		"\U00020000": "0",
		"\U0002b81d": "0",

		// NOTE: the CJK ranges in isAlphabet are hard coded and stop at
		// the boundaries above, so CJK runes which were assigned later
		// (Unicode 8.0 added U+9FCD..) are counted as words.
		"\u4d86": "1",
		"\u9fcd": "1",
	})
}

func TestWidth(t *testing.T) {
	runner := func(str string) string {
		return fmt.Sprint(Width(str))
	}

	runTestCases(t, runner, _M{
		"abcd\t0123\n7890": "12",
		"中zh英eng文混排":       "15",
		"":                 "0",

		// Control runes have zero width.
		"a\x00b":  "2",
		"a\nb\tc": "3",

		"中":       "2",
		"，。":      "4",
		"😀":       "2",
		"\u0301a": "2",

		// U+FFFD is the same rune as the one used to mark invalid bytes,
		// therefore its width is zero as well.
		"\uFFFD": "0",
	})
}

func TestRuneWidth(t *testing.T) {
	runner := func(str string) string {
		return fmt.Sprint(RuneWidth([]rune(str)[0]))
	}

	runTestCases(t, runner, _M{
		"a":    "1",
		"中":    "2",
		"\x11": "0",

		// Boundaries of the width table.
		"\x00":   "0",
		"\x1f":   "0",
		"\x20":   "1",
		"\x7f":   "1",
		"\u00a0": "1",
		"\u1fff": "1",
		"\u2000": "2",
		"\u3000": "2",
		"\uff60": "2",
		"\uff61": "1",
		"\uff9f": "1",
		"\uffa0": "2",
		"\uffe6": "2",
		"😀":      "2",
		"\u0301": "1",

		// RuneError is treated as a zero width rune.
		"\uFFFD": "0",

		// NOTE: PHP's mb_strwidth, which is referenced by RuneWidth's
		// document, reports 0 for "\x7f" and 1 for "\u2018". The values
		// asserted here are the current xstrings behaviour.
		"\u2018": "2",
	})
}
