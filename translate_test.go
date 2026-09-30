// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestTranslate(t *testing.T) {
	runner := func(str string) string {
		input := strings.Split(str, separator)
		return Translate(input[0], input[1], input[2])
	}

	runTestCases(t, runner, _M{
		sep("hello", "aeiou", "12345"):    "h2ll4",
		sep("hello", "aeiou", ""):         "hll",
		sep("hello", "a-z", "A-Z"):        "HELLO",
		sep("hello", "z-a", "a-z"):        "svool",
		sep("hello", "aeiou", "*"):        "h*ll*",
		sep("hello", "^l", "*"):           "**ll*",
		sep("hello", "p-z", "*"):          "hello",
		sep("hello ^ world", `\^lo`, "*"): "he*** * w*r*d",

		sep("中文字符测试", "文中谁敢试？", "123456"):  "21字符测5",
		sep("中文字符测试", "^文中谁敢试？", "123456"): "中文666试",
		sep("中文字符测试", "字-试", "0-9"):        "中90999",

		sep("h1e2l3l4o, w5o6r7l8d", "a-z,0-9", `A-Z\-a-czk-p`):       "HbEcLzLkO- WlOmRnLoD",
		sep("h1e2l3l4o, w5o6r7l8d", "a-zoh-n", "b-zakt-z"):           "t1f2x3x4k, x5k6s7x8e",
		sep("h1e2l3l4o, w5o6r7l8d", "helloa-zoh-n", "99999b-zakt-z"): "t1f2x3x4k, x5k6s7x8e",

		sep("hello", "e-", "p"):        "hpllo",
		sep("hello", "-e-", "p"):       "hpllo",
		sep("hello", "----e---", "p"):  "hpllo",
		sep("hello", "^---e----", "p"): "peppp",

		sep("hel\uFFFDlo", "\uFFFD", "H"): "helHlo",
		// A reverted pattern matches every rune but U+FFFD, so U+FFFD is
		// kept as is instead of being dropped. See #65.
		sep("hel\uFFFDlo", "^\uFFFD", "H"):   "HHH\uFFFDHH",
		sep("hel\uFFFDlo", "o-\uFFFDh", "H"): "HelHlH",

		// An empty from pattern means nothing is translated.
		sep("hello", "", "x"): "hello",
		sep("hello", "", ""):  "hello",

		sep("hello", "h", "H"):     "Hello",
		sep("a\x00b", "\x00", "X"): "aXb",

		// If no rune matches, the input is returned untouched.
		sep("a\xffb", "z", "X"): "a\xffb",

		// "^" alone matches every rune.
		sep("abc", "^", "x"): "xxx",
		sep("abc", "^", ""):  "",

		// If to is shorter than from, the last rune in to is repeated.
		sep("abcdef", "a-f", "xy"): "xyyyyy",
		sep("abc", "a-c", "x-y"):   "xyy",

		// Multiple from ranges are mapped one by one.
		sep("abcdef", "a-cx-z", "1-34-5"): "123def",

		// A range can be mapped to a single rune.
		sep("hello", "a-k", "λ"): "λλllo",

		// A reverted pattern uses the last rune of the to pattern only.
		sep("hello", "^e-l", "1-3"): "hell3",

		// A range may go backwards.
		sep("zyx", "z-x", "a-c"): "abc",
		sep("abc", "c-a", "1-3"): "321",
		sep("abc", "abc", "z-a"): "zyx",

		// A single rune in from is mapped to the first rune in to.
		sep("hello", "l", "123"): "he11o",

		// A range whose start and end are the same is a single rune.
		sep("abc", "a-a", "x"): "xbc",

		// A leading or trailing "-" is ignored.
		sep("a-b", "a-", "x"): "x-b",
		sep("a-b", "-a", "x"): "x-b",

		// "-" must be escaped to be used as a normal character.
		sep("a-b", `\-`, "x"): "axb",

		// "\\" matches a single backslash.
		sep(`a\b`, `\\`, "x"): "axb",

		// An empty to pattern deletes matched runes.
		sep("a-c", "a-c", ""): "-",

		// A rune range which overlaps a single rune recorded earlier
		// takes over that rune.
		sep("中", "中一-龥", "XY"):  "Y",
		sep("中一", "中一-龥", "XY"): "YY",

		// A rune which is not matched is kept as is, including a valid
		// U+FFFD rune and invalid bytes. See #65.
		sep("a\uFFFDb", "a", "x"):      "x\uFFFDb",
		sep("\uFFFDa", "a", "x"):       "\uFFFDx",
		sep("a\xffb", "a", "x"):        "x\xffb",
		sep("a\xff\xfe b", "a-b", "X"): "X\xff\xfe X",

		// A rune can be translated to a NUL rune. See #66.
		sep("hello", "h", "\x00"):  "\x00ello",
		sep("hello", "l", "\x00x"): "he\x00\x00o",
	})
}

func TestTranslateReplacementCharacterPatterns(t *testing.T) {
	cases := []struct {
		name, str, from, to, want string
	}{
		{"from start", "\uFFFDa", "\uFFFDa", "xy", "xy"},
		{"from middle", "a\uFFFDb", "a\uFFFDb", "xyz", "xyz"},
		{"from end", "a\uFFFD", "a\uFFFD", "xy", "xy"},
		{"to start", "ab", "ab", "\uFFFDx", "\uFFFDx"},
		{"to middle", "abc", "abc", "x\uFFFDy", "x\uFFFDy"},
		{"to end", "ab", "ab", "x\uFFFD", "x\uFFFD"},
		{"repeated last rune", "abc", "abc", "x\uFFFD", "x\uFFFD\uFFFD"},
		{"from ascending range", "\uFFFC\uFFFD", "\uFFFC-\uFFFD", "a-b", "ab"},
		{"from descending range", "\uFFFD\uFFFC", "\uFFFD-\uFFFC", "a-b", "ab"},
		{"to ascending range", "ab", "ab", "\uFFFC-\uFFFD", "\uFFFC\uFFFD"},
		{"to descending range", "ab", "ab", "\uFFFD-\uFFFC", "\uFFFD\uFFFC"},
		{"escaped literal", "\uFFFDa", "\\\uFFFDa", "xy", "xy"},
		{"reverted pattern", "\uFFFDab", "^\uFFFDa", "x", "\uFFFDax"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Translate(c.str, c.from, c.to); got != c.want {
				t.Errorf("Translate(%q, %q, %q) = %q, want %q", c.str, c.from, c.to, got, c.want)
			}
		})
	}
}

func TestTranslatorIgnoredPatternCharacters(t *testing.T) {
	// Keep the existing replacement-rune fallback for patterns with no literals.
	for _, pattern := range []string{"-", "---", "\\"} {
		if got := Translate("\uFFFDab", pattern, "x"); got != "xab" {
			t.Errorf("Translate with from %q = %q, want %q", pattern, got, "xab")
		}
		if got := Translate("\uFFFDab", "^"+pattern, "x"); got != "\uFFFDxx" {
			t.Errorf("Translate with reverted from %q = %q", pattern, got)
		}
		tr := NewTranslator("ab", pattern)
		if got, matched := tr.TranslateRune('a'); !matched || got != utf8.RuneError {
			t.Errorf("TranslateRune with to %q = (%U, %v)", pattern, got, matched)
		}
	}
}

func TestDeleteCountReplacementCharacterPatterns(t *testing.T) {
	str := "\uFFFDa\uFFFDb"
	for _, pattern := range []string{"\uFFFDa", "a\uFFFD"} {
		if got := Delete(str, pattern); got != "b" {
			t.Errorf("Delete(%q, %q) = %q, want %q", str, pattern, got, "b")
		}
		if got := Count(str, pattern); got != 3 {
			t.Errorf("Count(%q, %q) = %d, want 3", str, pattern, got)
		}
		tr := NewTranslator(pattern, "")
		if got, matched := tr.TranslateRune(utf8.RuneError); !matched || got != utf8.RuneError {
			t.Errorf("TranslateRune for deletion with from %q = (%U, %v)", pattern, got, matched)
		}
	}
}

func TestTranslateCountDeleteConsistency(t *testing.T) {
	// Count reports how many runes match the pattern and Delete removes
	// exactly those runes, so Delete must not drop anything else. This used
	// to fail for U+FFFD and for invalid bytes. See #65.
	inputs := []string{"a\uFFFDb", "a\xffb", "\uFFFD\uFFFD", "hello 世界", "中\uFFFDb", ""}
	patterns := []string{"a-z", "a", "^a-z", "中", "\uFFFD", "^"}

	for _, str := range inputs {
		for _, pattern := range patterns {
			got := Len(Delete(str, pattern)) + Count(str, pattern)

			if want := Len(str); got != want {
				t.Fatalf("Len(Delete(%q, %q)) + Count(%q, %q) = %d, want %d",
					str, pattern, str, pattern, got, want)
			}
		}
	}
}

func TestDelete(t *testing.T) {
	runner := func(str string) string {
		input := strings.Split(str, separator)
		return Delete(input[0], input[1])
	}

	runTestCases(t, runner, _M{
		sep("hello", "aeiou"): "hll",
		sep("hello", "a-k"):   "llo",
		sep("hello", "^a-k"):  "he",

		sep("中文字符测试", "文中谁敢试？"): "字符测",

		// An empty pattern deletes nothing.
		sep("hello", ""): "hello",
		sep("", "a"):     "",

		// "^" alone deletes every rune.
		sep("abc", "^"): "",

		sep("a\x00b", "\x00"): "ab",

		// Nothing matches, the input is returned untouched.
		sep("a\uFFFDb", "z"): "a\uFFFDb",
	})
}

func TestCount(t *testing.T) {
	runner := func(str string) string {
		input := strings.Split(str, separator)
		return fmt.Sprint(Count(input[0], input[1]))
	}

	runTestCases(t, runner, _M{
		sep("hello", "aeiou"): "2",
		sep("hello", "a-k"):   "2",
		sep("hello", "^a-k"):  "3",

		sep("中文字符测试", "文中谁敢试？"): "3",

		// An empty string or pattern counts zero runes.
		sep("hello", ""): "0",
		sep("", "a"):     "0",

		sep("hello", "l"):       "2",
		sep("hello", "^e-l"):    "1",
		sep("abc", "^"):         "3",
		sep("abcdef", "a-cx-z"): "3",
	})
}

func TestSqueeze(t *testing.T) {
	runner := func(str string) string {
		input := strings.Split(str, separator)
		return Squeeze(input[0], input[1])
	}

	runTestCases(t, runner, _M{
		sep("hello", ""):             "helo",
		sep("hello     world", ""):   "helo world",
		sep("hello     world", " "):  "hello world",
		sep("hello     world", "  "): "hello world",
		sep("hello", "a-k"):          "hello",
		sep("hello", "^a-k"):         "helo",
		sep("hello", "^a-l"):         "hello",
		sep("foooo baaaaar", "a"):    "foooo bar",

		sep("打打打打个劫！！", ""):  "打个劫！",
		sep("打打打打个劫！！", "打"): "打个劫！！",

		sep("", ""):          "",
		sep("a", ""):         "a",
		sep("aa", ""):        "a",
		sep("aaa", "b"):      "aaa",
		sep("aaabbb", "a-b"): "ab",
		sep("中中中", "中"):      "中",

		// "-" is a pattern character, escape it to match a literal dash.
		sep("a-b--c", `\-`): "a-b-c",
	})
}

func TestTranslatorHasPattern(t *testing.T) {
	cases := []struct {
		from string
		to   string
		want bool
	}{
		{"", "", false},
		{"", "x", false},
		{"a", "", true},
		{"a", "b", true},
		{"^", "", true},
		{"^a", "b", true},
	}

	for _, c := range cases {
		if got := NewTranslator(c.from, c.to).HasPattern(); got != c.want {
			t.Fatalf("NewTranslator(%q, %q).HasPattern() = %v, want %v", c.from, c.to, got, c.want)
		}
	}
}

func TestTranslatorTranslateRune(t *testing.T) {
	cases := []struct {
		from, to string
		r        rune
		result   rune
		matched  bool
	}{
		// A regular 1:1 mapping.
		{"a-c", "x-z", 'a', 'x', true},
		{"a-c", "x-z", 'b', 'y', true},
		{"a-c", "x-z", 'c', 'z', true},
		{"a-c", "x-z", 'd', 'd', false},
		{"a-c", "x-z", '中', '中', false},

		// A reverted pattern matches every rune but the listed ones.
		{"^a-c", "z", 'a', 'a', false},
		{"^a-c", "z", 'd', 'z', true},

		// In "delete" mode (empty to pattern) matched runes are mapped to
		// utf8.RuneError, which is then dropped by Translate.
		{"ab", "", 'a', utf8.RuneError, true},
		{"ab", "", 'x', 'x', false},

		// A mapping whose target rune is U+0000 is a regular mapping. See #66.
		{"h", "\x00", 'h', 0, true},
		{"h", "\x00", 'x', 'x', false},

		// Without any pattern every rune is returned as is.
		{"", "", 'a', 'a', false},
	}

	for _, c := range cases {
		tr := NewTranslator(c.from, c.to)
		result, matched := tr.TranslateRune(c.r)

		if result != c.result || matched != c.matched {
			t.Fatalf("NewTranslator(%q, %q).TranslateRune(%q) = (%q, %v), want (%q, %v)",
				c.from, c.to, c.r, result, matched, c.result, c.matched)
		}
	}
}

func TestTranslatorReuse(t *testing.T) {
	tr := NewTranslator("aeiou", "12345")

	runTestCases(t, tr.Translate, _M{
		"hello": "h2ll4",
		"world": "w4rld",
		"":      "",
		"aeiou": "12345",
		"xyz":   "xyz",
	})
}
