package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Go(func() {
		for {
			select {
			case <-done:
				return
			case job, ok := <-jobs:
				if !ok {
					return
				}
				fmt.Printf("processing job %d\n", job)
			}
		}
	})

	jobs <- 1
	jobs <- 2

	close(done)
	wg.Wait()
	fmt.Println("main ends")
}
