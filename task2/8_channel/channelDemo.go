package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 10)
	var wg sync.WaitGroup

	//producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}

		close(ch)
	}()

	//consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("receive %d \n", num)
		}
	}()

	wg.Wait()
	fmt.Printf("finish two goroutine task")
}
