package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
