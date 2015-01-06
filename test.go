// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"testing"
)

func runTestCases(t *testing.T, converter func(string)string, cases map[string]string) {
	for k, v := range cases {
		s := converter(k)

		if s != v {
			t.Fatalf("case fails. [expected:%v] [actual:%v]", v, s)
		}
	}
}