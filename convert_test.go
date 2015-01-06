// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"testing"
)

func TestToUnderscore(t *testing.T) {
	runTestCases(t, ToUnderscore, map[string]string{
		"HTTPServer": "http_server",
		"_camelCase": "_camel_case",
		"NoHTTPS": "no_https",
		"Wi_thF":             "wi_th_f",
		"_AnotherTES_TCaseP": "__another_te_s__t_case_p",
		"ALL": "all",
	})
}

func TestToCamelCase(t *testing.T) {
	runTestCases(t, ToCamelCase, map[string]string{
		"http_server": "HttpServer",
		"_camel_case": "_CamelCase",
		"no_https": "NoHttps",
		"_complex__case_":             "_Complex_Case_",
		"all": "All",
	})
}

