# mutable

[![GoDoc](https://godoc.org/github.com/notduncansmith/mutable?status.svg)](https://godoc.org/github.com/notduncansmith/mutable)

`mutable.RW` wraps a [`sync.RWMutex`](https://golang.org/pkg/sync/#RWMutex) with functions for reading/writing it. It is intended to be embedded within structs that require cross-thread synchronization.

## Usage

See GoDoc for per-method docs

```go
// example/main.go
package main

import (
	"log"

	mutable "github.com/notduncansmith/mutable"
)

type counter struct {
	*mutable.RW
	i int
}

func (c *counter) inc() int {
	return c.WithRWLock(func() interface{} {
		c.i++
		return c.i
	}).(int)
}

func (c *counter) v() int {
	return c.WithRLock(func() interface{} {
		return c.i
	}).(int)
}

func (c *counter) reset() {
	c.DoWithRWLock(func() {
		c.i = 0
	})
}

func (c *counter) print() {
	// technically unnecessary given that v() acquires lock
	c.DoWithRLock(func() {
		log.Printf("Count: %v", c.i)
	})
}

func main() {
	c := counter{mutable.NewRW("thing"), 0}
	c.print()
	c.inc()
	c.print()
	c.inc()
	c.print()
	c.reset()
	c.print()
}

// ❯ go build . && ./example
// 2019/11/04 23:22:22 Count: 0
// 2019/11/04 23:22:22 Count: 1
// 2019/11/04 23:22:22 Count: 2
// 2019/11/04 23:22:22 Count: 0
```

## License

[The MIT License](https://opensource.org/licenses/MIT):

Copyright 2019 Duncan Smith

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.