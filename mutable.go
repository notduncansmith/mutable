package mutable

import (
	"sync"
)

// RW wraps a `sync.RWMutex` with functions for reading/writing it. It is intended to be embedded within structs that require cross-thread synchronization
type RW struct {
	mut  *sync.RWMutex
	Name string
}

// NewRW returns a new RW
func NewRW(name string) *RW {
	mut := sync.RWMutex{}
	return &RW{&mut, name}
}

// WithRLock (f) calls `f` while holding a Read Lock on `mut`
func (rw *RW) WithRLock(f func() interface{}) interface{} {
	rw.mut.RLock()
	defer rw.mut.RUnlock()
	result := f()
	return result
}

// WithRWLock (f) calls `f` while holding a Read/Write Lock on `mut`
func (rw *RW) WithRWLock(f func() interface{}) interface{} {
	rw.mut.Lock()
	defer rw.mut.Unlock()
	result := f()
	return result
}

// DoWithRLock (f) is like WithRLock but does not return a value
func (rw *RW) DoWithRLock(f func()) {
	rw.mut.RLock()
	defer rw.mut.RUnlock()
	f()
}

// DoWithRWLock (f) is like WithLock but does not return a value
func (rw *RW) DoWithRWLock(f func()) {
	rw.mut.Lock()
	defer rw.mut.Unlock()
	f()
}
