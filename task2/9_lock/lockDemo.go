package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	const (
		goroutines = 10
		iter       = 1000
	)

	wg.Add(goroutines)
	for i := 1; i <= goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 1; j <= iter; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("finish counter plus: %d\n", counter)
}
