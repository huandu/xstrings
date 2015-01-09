// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	LenRunner := func(str string) string {
		return fmt.Sprint(Len(str))
	}

	runTestCases(t, LenRunner, _M{
		"abcdef":    "6",
		"中文":        "2",
		"中yin文hun排": "9",
		"":          "0",
	})
}
