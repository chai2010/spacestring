// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// Package spacestring implements simple functions to manipulate spacestring.
package spacestring

import "regexp"

var re = regexp.MustCompile(`^\s+$`)

// Space is a spacestring.
const Space = " "

// IsSpace reports whether its argument s is spacestring.
func IsSpace(s string) bool {
	return re.MatchString(s)
}
