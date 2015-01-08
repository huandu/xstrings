// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"testing"
)

func TestReverse(t *testing.T) {
	runTestCases(t, Reverse, map[string]string{
		"reverse string": "gnirts esrever",
		"中文如何？":          "？何如文中",
		"中en文混~排怎样？a":    "a？样怎排~混文ne中",
	})
}
