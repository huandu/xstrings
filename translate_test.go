// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"strings"
	"testing"
)

func TestTranslate(t *testing.T) {
	runner := func(str string) string {
		inputs := strings.Split(str, separator)
		return Translate(inputs[0], inputs[1], inputs[2])
	}

	runTestCases(t, runner, _M{
		sep("hello", "aeiou", "12345"):    "h2ll4",
		sep("hello", "a-z", "A-Z"):        "HELLO",
		sep("hello", "z-a", "a-z"):        "svool",
		sep("hello", "aeiou", "*"):        "h*ll*",
		sep("hello", "^l", "*"):           "**ll*",
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
	})
}
