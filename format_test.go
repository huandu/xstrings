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

		// A tab is expanded to at least one space, even at column 0.
		sep("\t", "4"):   "    ",
		sep("\t", "1"):   " ",
		sep("\t\t", "3"): "      ",

		// Trailing tabs and tabs after a newline are expanded as usual.
		sep("a\t", "4"):    "a   ",
		sep("\n\t", "4"):   "\n    ",
		sep("\n\n\t", "2"): "\n\n  ",

		// A wide rune occupies two columns.
		sep("中\t", "4"):   "中  ",
		sep("a中\tb", "4"): "a中 b",

		sep("x\ty\tz", "8"): "x       y       z",

		// NOTE: only '\n' resets the column, as documented. Python's
		// str.expandtabs resets it on '\r' too.
		sep("a\rb\tc", "4"): "a\rb  c",

		sep("abc", "0"):  "tab size must be positive",
		sep("abc", "-1"): "tab size must be positive",
	})
}

func TestExpandTabsLongString(t *testing.T) {
	// A long input makes allocBuffer clamp its initial capacity.
	str := strings.Repeat("a", 600) + "\t" + strings.Repeat("中", 600) + "\t"
	want := strings.Repeat("a", 600) + "    " + strings.Repeat("中", 600) + "    "

	if got := ExpandTabs(str, 4); got != want {
		t.Fatalf("ExpandTabs(long string, 4) = %q, want %q", got, want)
	}
}

func TestLeftJustify(t *testing.T) {
	runner := func(str string) string {
		input := strings.Split(str, separator)
		n, _ := strconv.Atoi(input[1])
		return LeftJustify(input[0], n, input[2])
	}

	runTestCases(t, runner, _M{
		sep("hello", "4", " "):    "hello",
		sep("hello", "10", " "):   "hello     ",
		sep("hello", "10", "123"): "hello12312",

		sep("hello中文test", "4", " "):    "hello中文test",
		sep("hello中文test", "12", " "):   "hello中文test ",
		sep("hello中文test", "18", "测试！"): "hello中文test测试！测试！测",

		sep("hello中文test", "0", "123"): "hello中文test",
		sep("hello中文test", "18", ""):   "hello中文test",

		// Length is measured in runes, pad in runes as well.
		sep("", "5", "ab"):    "ababa",
		sep("abc", "3", "x"):  "abc",
		sep("abc", "-1", "x"): "abc",
		sep("中文", "4", "-"):   "中文--",
		sep("中文", "4", "中文"):  "中文中文",
		sep("a", "5", "中文"):   "a中文中文",
		sep("a", "3", "中文"):   "a中文",

		// A pad string is cut in the middle of its runes when needed.
		sep("a", "2", "中文"): "a中",
	})
}

func TestRightJustify(t *testing.T) {
	runner := func(str string) string {
		input := strings.Split(str, separator)
		n, _ := strconv.Atoi(input[1])
		return RightJustify(input[0], n, input[2])
	}

	runTestCases(t, runner, _M{
		sep("hello", "4", " "):    "hello",
		sep("hello", "10", " "):   "     hello",
		sep("hello", "10", "123"): "12312hello",

		sep("hello中文test", "4", " "):    "hello中文test",
		sep("hello中文test", "12", " "):   " hello中文test",
		sep("hello中文test", "18", "测试！"): "测试！测试！测hello中文test",

		sep("hello中文test", "0", "123"): "hello中文test",
		sep("hello中文test", "18", ""):   "hello中文test",

		sep("", "5", "ab"):   "ababa",
		sep("abc", "3", "x"): "abc",
		sep("中文", "4", "-"):  "--中文",
		sep("中文", "4", "中文"): "中文中文",
		sep("a", "5", "中文"):  "中文中文a",
		sep("a", "3", "中文"):  "中文a",
		sep("a", "2", "中文"):  "中a",
	})
}

func TestCenter(t *testing.T) {
	runner := func(str string) string {
		input := strings.Split(str, separator)
		n, _ := strconv.Atoi(input[1])
		return Center(input[0], n, input[2])
	}

	runTestCases(t, runner, _M{
		sep("hello", "4", " "):    "hello",
		sep("hello", "10", " "):   "  hello   ",
		sep("hello", "10", "123"): "12hello123",

		sep("hello中文test", "4", " "):    "hello中文test",
		sep("hello中文test", "12", " "):   "hello中文test ",
		sep("hello中文test", "18", "测试！"): "测试！hello中文test测试！测",

		sep("hello中文test", "0", "123"): "hello中文test",
		sep("hello中文test", "18", ""):   "hello中文test",

		sep("", "5", "ab"):   "ababa",
		sep("abc", "3", "x"): "abc",
		sep("中文", "4", "-"):  "-中文-",
		sep("中文", "4", "中文"): "中中文中",
		sep("a", "5", "中文"):  "中文a中文",
		sep("a", "3", "中文"):  "中a中",
		sep("a", "2", "中文"):  "a中",

		// No padding is added when there is no room for it.
		sep("x", "0", "y"): "x",
		sep("x", "1", "y"): "x",
	})
}
