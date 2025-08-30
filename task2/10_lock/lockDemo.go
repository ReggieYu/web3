package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		counter int64
		wg      sync.WaitGroup
	)

	const (
		goroutines = 10
		iter       = 1000
	)

	wg.Add(goroutines)
	for i := 1; i <= goroutines; i++ {
		go func() {
			defer wg.Done()
			for k := 1; k <= iter; k++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("finish counter plus: %d\n", counter)
}
