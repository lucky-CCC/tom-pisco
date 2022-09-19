package pslice

import (
	"sync"
	"sync/atomic"
	"testing"
)

// go test -v -race -run Free
// 会出现DATA RACE
func TestFree(t *testing.T) {
	slice := make([]int, 0)
	wg := sync.WaitGroup{}
	num := 10
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			slice = append(slice, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(slice)
}

// go test -v -race -run Lock
// 加了锁，运行正常
func TestLock(t *testing.T) {
	slice := make([]int, 0)
	wg := sync.WaitGroup{}
	num := 10
	wg.Add(num)
	lock := sync.Mutex{}

	for i := 0; i < num; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			slice = append(slice, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(slice)
}

// go test -v -race -run CasLock
// 加了锁，运行正常
func TestCasLock(t *testing.T) {
	slice := make([]int, 0)
	wg := sync.WaitGroup{}
	num := 10
	wg.Add(num)
	casLock := atomic.Bool{}
	casLock.Store(false)

	for i := 0; i < num; i++ {
		go func() {
			for {
				if casLock.CompareAndSwap(false, true) {
					slice = append(slice, 1)
					wg.Done()
					casLock.Store(false)
					break
				}
			}
		}()
	}
	wg.Wait()
	t.Log(slice)
}

// go test -bench BenchmarkLock -benchmem
func BenchmarkLock(b *testing.B) {
	slice := make([]int, 0)
	wg := sync.WaitGroup{}
	num := b.N
	wg.Add(num)
	lock := sync.Mutex{}

	for i := 0; i < num; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			slice = append(slice, 1)
			wg.Done()
		}()
	}
	wg.Wait()
}

// go test -bench BenchmarkCasLock -benchmem -benchtime=10000x
func BenchmarkCasLock(b *testing.B) {
	slice := make([]int, 0)
	wg := sync.WaitGroup{}
	num := b.N
	wg.Add(num)
	casLock := atomic.Bool{}
	casLock.Store(false)

	for i := 0; i < num; i++ {
		go func() {
			for {
				if casLock.CompareAndSwap(false, true) {
					slice = append(slice, 1)
					wg.Done()
					casLock.Store(false)
					break
				}
			}
		}()
	}
	wg.Wait()
}
