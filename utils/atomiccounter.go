package utils

import (
	"sync/atomic"
)

// AtomicCounter is an abstraction of a simple count variable
// It implements methods to add atomically and to fetch the current count
type AtomicCounter struct {
	count uint64
}

// Add adds 1 to count
func (ac *AtomicCounter) Add() {
	atomic.AddUint64(&ac.count, 1)
}

// GetCount returns the count on the AtomicCounter type
func (ac *AtomicCounter) GetCount() int {
	return int(ac.count)
}
