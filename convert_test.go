// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	runTestCases(t, ToSnakeCase, map[string]string{
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

		"  sentence case  ":                                    "__sentence_case__",
		" Mixed-hyphen case _and SENTENCE_case and UPPER-case": "_mixed_hyphen_case__and_sentence_case_and_upper_case",
	})
}

func TestToCamelCase(t *testing.T) {
	runTestCases(t, ToCamelCase, map[string]string{
		"http_server":     "HttpServer",
		"_camel_case":     "_CamelCase",
		"no_https":        "NoHttps",
		"_complex__case_": "_Complex_Case_",
		"all":             "All",
	})
}

func TestSwapCase(t *testing.T) {
	runTestCases(t, SwapCase, map[string]string{
		"swapCase": "SWAPcASE",
		"Θ~λa云Ξπ":  "θ~ΛA云ξΠ",
	})
}

func TestFirstRuneToUpper(t *testing.T) {
	runTestCases(t, FirstRuneToUpper, map[string]string{
		"hello, world!": "Hello, world!",
		"Hello, world!": "Hello, world!",
		"你好，世界！":        "你好，世界！",
	})
}

func TestFirstRuneToLower(t *testing.T) {
	runTestCases(t, FirstRuneToLower, map[string]string{
		"hello, world!": "hello, world!",
		"Hello, world!": "hello, world!",
		"你好，世界！":        "你好，世界！",
	})
}
