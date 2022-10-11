package main

import (
	"sync"
	"testing"
)

func TestWg(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	group.Done()
	group.Wait()
}
