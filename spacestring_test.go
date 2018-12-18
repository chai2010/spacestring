// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package spacestring_test

import (
	"testing"

	"github.com/chai2010/spacestring"
)

func TestIsSpace(t *testing.T) {
	for i, v := range []struct {
		s  string
		ok bool
	}{
		// space string
		{spacestring.Space, true},
		{" ", true},
		{" \t\n\r", true},
		{" \t\n\r ", true},

		// non space string
		{"", false},
		{" abc ", false},
		{" 中国·武汉 ", false},
	} {
		if spacestring.IsSpace(v.s) != v.ok {
			t.Fatalf("%d: %q failed\n", i, v.s)
		}
	}
}
