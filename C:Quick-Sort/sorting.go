package main

import (
	"time"
)

type ByDuration []time.Duration

// sorting
func (a ByDuration) Len() int           { return len(a) }
func (a ByDuration) Less(i, j int) bool { return a[i].Nanoseconds() < a[j].Nanoseconds() }
func (a ByDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
