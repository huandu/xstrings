// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"sort"
	"strings"
	"testing"
)

func TestToSnakeCaseAndToKebabCase(t *testing.T) {
	cases := _M{
		"HTTPServer":         "http_server",
		"_camelCase":         "_camel_case",
		"NoHTTPS":            "no_https",
		"Wi_thF":             "wi_th_f",
		"_AnotherTES_TCaseP": "_another_tes_t_case_p",
		"ALL":                "all",
		"_HELLO_WORLD_":      "_hello_world_",
		"HELLO_WORLD":        "hello_world",
		"HELLO____WORLD":     "hello____world",
		"TW":                 "tw",
		"_C":                 "_c",
		"http2xx":            "http_2xx",
		"HTTP2XX":            "http2_xx",
		"HTTP20xOK":          "http_20x_ok",
		"HTTP20xStatus":      "http_20x_status",
		"HTTP-20xStatus":     "http_20x_status",
		"a":                  "a",
		"Duration2m3s":       "duration_2m3s",
		"Bld4Floor3rd":       "bld4_floor_3rd",
		" _-_ ":              "_____",
		"a1b2c3d":            "a_1b2c3d",
		"A//B%%2c":           "a//b%%2c",

		"HTTP状态码404/502Error": "http_状态码404/502_error",
		"中文(字符)":              "中文(字符)",
		"混合ABCWords与123数字456": "混合_abc_words_与123_数字456",

		"  sentence case  ": "__sentence_case__",
		" Mixed-hyphen case _and SENTENCE_case and UPPER-case": "_mixed_hyphen_case__and_sentence_case_and_upper_case",
		"FROM CamelCase to snake/kebab-case":                   "from_camel_case_to_snake/kebab_case",

		"": "",
		"Abc\uFFFDE\uFFFDf\uFFFDd\uFFFD2\uFFFD00z\uFFFDZZ\uFFFDZZ": "abc_\uFFFDe\uFFFDf\uFFFDd_\uFFFD2\uFFFD00z_\uFFFDzz\uFFFDzz",
		"\uFFFD\uFFFD\uFFFD\uFFFD\uFFFD":                           "\uFFFD\uFFFD\uFFFD\uFFFD\uFFFD",

		"abc_123_def": "abc_123_def",

		// Samples from the package document.
		"FirstName": "first_name",
		"GO_PATH":   "go_path",
		"GO PATH":   "go_path",
		"GO-PATH":   "go_path",

		// Already separated words are kept.
		"HTTP_20x_OK": "http_20x_ok",
		"1_2":         "1_2",

		"aB":   "a_b",
		"AB":   "ab",
		"Ab":   "ab",
		"A1":   "a1",
		"AB12": "ab12",

		// A digit in front of a word is kept in the same word, a digit
		// before an upper case rune gets its own connector.
		"12AB": "12_ab",
		"1A":   "1_a",

		// Punctuation and non-alphabet runes are kept as is.
		"a.b":  "a.b",
		"a'b":  "a'b",
		"a😀b":  "a_😀_b",
		"😀ABC": "😀_abc",

		// CJK runes are not alphabetic, so they act like other runes.
		"中ABC": "中_abc",
		"ABC中": "abc_中",
	}

	runTestCases(t, ToSnakeCase, cases)

	for k, v := range cases {
		cases[k] = strings.Replace(v, "_", "-", -1)
	}

	runTestCases(t, ToKebabCase, cases)
}

func TestToCamelCase(t *testing.T) {
	runTestCases(t, ToCamelCase, _M{
		"http_server":     "httpServer",
		"_camel_case":     "_camelCase",
		"no_https":        "noHttps",
		"_complex__case_": "_complex_Case_",
		" complex -case ": " complex Case ",
		"all":             "all",
		"GOLANG_IS_GREAT": "golangIsGreat",
		"GOLANG":          "golang",
		"a":               "a",
		"好":               "好",

		"FROM CamelCase to snake/kebab-case": "fromCamelCaseToSnake/kebabCase",

		"": "",

		// Samples from the package document.
		"some words":   "someWords",
		"hello-world":  "helloWorld",
		"hello world":  "helloWorld",
		"HELLO":        "hello",
		"helloWORLD":   "helloWORLD",
		"some_1_thing": "some1Thing",
		"1_abc":        "1Abc",
		"abc_1":        "abc1",
		"v1_2v":        "v12v",
		"a__b":         "a_B",
		"a  b":         "a B",
		"1_2":          "12",
		"_a_B_":        "_aB_",
		"中文_测试":        "中文测试",
	})
}

func TestToPascalCase(t *testing.T) {
	runTestCases(t, ToPascalCase, _M{
		"http_server":     "HttpServer",
		"_camel_case":     "_CamelCase",
		"no_https":        "NoHttps",
		"_complex__case_": "_Complex_Case_",
		" complex -case ": " Complex Case ",
		"all":             "All",
		"GOLANG_IS_GREAT": "GolangIsGreat",
		"GOLANG":          "Golang",
		"a":               "A",
		"好":               "好",

		"FROM CamelCase to snake/kebab-case": "FromCamelCaseToSnake/kebabCase",

		"": "",

		"some words":   "SomeWords",
		"hello-world":  "HelloWorld",
		"HELLO":        "Hello",
		"helloWORLD":   "HelloWORLD",
		"some_1_thing": "Some1Thing",
		"1_abc":        "1Abc",
		"abc_1":        "Abc1",
		"v1_2v":        "V12v",
		"a__b":         "A_B",
		"a  b":         "A B",
		"1_2":          "12",
		"_a_B_":        "_AB_",
		"中文_测试":        "中文测试",
	})
}

func TestSwapCase(t *testing.T) {
	runTestCases(t, SwapCase, _M{
		"swapCase": "SWAPcASE",
		"Θ~λa云Ξπ":  "θ~ΛA云ξΠ",
		"a":        "A",

		"": "",

		"A":     "a",
		"HELLO": "hello",
		"Hello": "hELLO",
		"1":     "1",
		"中":     "中",
		"ß":     "ß", // unicode.ToUpper("ß") is "ß" in Go.
		"İ":     "i",
	})
}

func TestFirstRuneToUpper(t *testing.T) {
	runTestCases(t, FirstRuneToUpper, _M{
		"hello, world!": "Hello, world!",
		"Hello, world!": "Hello, world!",
		"你好，世界！":        "你好，世界！",
		"a":             "A",

		"": "",

		// A first rune which is not a lower case letter is kept as is.
		"1abc":    "1abc",
		"中abc":    "中abc",
		"ßabc":    "ßabc",
		"\xffabc": "\xffabc",
	})
}

func TestFirstRuneToLower(t *testing.T) {
	runTestCases(t, FirstRuneToLower, _M{
		"hello, world!": "hello, world!",
		"Hello, world!": "hello, world!",
		"你好，世界！":        "你好，世界！",
		"a":             "a",
		"A":             "a",

		"": "",

		"1abc": "1abc",
		"İabc": "iabc",
	})
}

func TestShuffle(t *testing.T) {
	// It seems there is no reliable way to test shuffled string.
	// Runner just make sure shuffled string has the same runes as origin string.
	runner := func(str string) string {
		s := Shuffle(str)
		slice := sort.StringSlice(strings.Split(s, ""))
		slice.Sort()
		return strings.Join(slice, "")
	}

	runTestCases(t, runner, _M{
		"":            "",
		"facgbheidjk": "abcdefghijk",
		"尝试中文":        "中尝文试",
		"zh英文hun排":    "hhnuz排文英",

		"a": "a",
		"中": "中",
	})
}

type testShuffleSource int

// A generated random number sequance just for testing.
var testShuffleTable = []int64{
	1874068156324778273,
	3328451335138149956,
	5263531936693774911,
	7955079406183515637,
	2703501726821866378,
	2740103009342231109,
	6941261091797652072,
	1905388747193831650,
	7981306761429961588,
	6426100070888298971,
	4831389563158288344,
	261049867304784443,
	1460320609597786623,
	5600924393587988459,
	8995016276575641803,
	732830328053361739,
	5486140987150761883,
	545291762129038907,
	6382800227808658932,
	2781055864473387780,
	1598098976185383115,
	4990765271833742716,
	5018949295715050020,
	2568779411109623071,
	3902890183311134652,
	4893789450120281907,
	2338498362660772719,
	2601737961087659062,
	7273596521315663110,
	3337066551442961397,
	8121576815539813105,
	2740376916591569721,
	8249030965139585917,
	898860202204764712,
	9010467728050264449,
	685213522303989579,
	2050257992909156333,
	6281838661429879825,
	2227583514184312746,
	2873287401706343734,
}

func (src *testShuffleSource) Int63() int64 {
	n := testShuffleTable[int(*src)%len(testShuffleTable)]
	(*src)++
	return n
}

func (*testShuffleSource) Seed(int64) {}

func TestShuffleSource(t *testing.T) {
	runner := func(str string) string {
		var src testShuffleSource
		return ShuffleSource(str, &src)
	}

	runTestCases(t, runner, _M{
		"":            "",
		"facgbheidjk": "bkgfijached",
		"尝试中文怎么样":     "怎试么中样尝文",
		"zh英文hun排":    "zuhh文n英排",

		"a": "a",
		"中": "中",
	})
}

func TestSuccessor(t *testing.T) {
	runTestCases(t, Successor, _M{
		"":          "",
		"abcd":      "abce",
		"THX1138":   "THX1139",
		"<<koala>>": "<<koalb>>",
		"1999zzz":   "2000aaa",
		"ZZZ9999":   "AAAA0000",
		"***":       "**+",

		"来点中文试试":               "来点中文试诖",
		"中cZ英ZZ文zZ混9zZ9杂99进z位": "中dA英AA文aA混0aA0杂00进a位",

		// Carries which are absorbed by an alphanumeric rune on the left.
		"z":  "aa",
		"zz": "aaa",
		"Z":  "AA",
		"y9": "z0",
		"a9": "b0",
		"Z9": "AA0",
		"z9": "aa0",
		"8":  "9",

		// A non-ASCII rune is incremented in place.
		"中":   "丮",
		"中中":  "中丮",
		"a中":  "b中",
		"a中z": "b中a",

		// Incrementing the largest rune wraps over and yields the
		// replacement rune, as documented.
		"\U0010FFFF": "\uFFFD",

		// A carry rune is inserted in front of the leftmost alphanumeric
		// rune, even when it is not the first rune of the string.
		"-zz": "-aaa",
		"-ZZ": "-AAA",
	})
}
