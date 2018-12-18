# spacestring

[![Build Status](https://travis-ci.org/chai2010/spacestring.svg)](https://travis-ci.org/chai2010/spacestring)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/spacestring)](https://goreportcard.com/report/github.com/chai2010/spacestring)
[![GoDoc](https://godoc.org/github.com/chai2010/spacestring?status.svg)](https://godoc.org/github.com/chai2010/spacestring)
[![License](http://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/chai2010/spacestring/blob/master/LICENSE)


- https://godoc.org/github.com/chai2010/spacestring
- spacestring contains more than 1 space characters
- empty string is not spacestring
- `^\s+$`is space string

## Example

```go
package main

import (
	"fmt"

	"github.com/chai2010/spacestring"
)

func main() {
	// nil string contains more than 1 space characters
	fmt.Println(spacestring.IsSpace(" "))
	fmt.Println(spacestring.IsSpace(" \t\n\r"))

	// empty string is not nil string
	fmt.Println(spacestring.IsSpace(""))

	// non space string is not nil string
	fmt.Println(spacestring.IsSpace(" abc "))

	// Output:
	// true
	// true
	// false
	// false
}
```

## BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
