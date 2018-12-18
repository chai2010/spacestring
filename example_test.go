// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package spacestring_test

import (
	"fmt"

	"github.com/chai2010/spacestring"
)

func Example() {
	// space string contains more than 1 space characters
	fmt.Println(spacestring.IsSpace(" "))
	fmt.Println(spacestring.IsSpace(" \t\n\r"))

	// empty string is not space string
	fmt.Println(spacestring.IsSpace(""))

	// non space string is not space string
	fmt.Println(spacestring.IsSpace(" abc "))

	// Output:
	// true
	// true
	// false
	// false
}
