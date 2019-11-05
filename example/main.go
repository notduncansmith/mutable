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
