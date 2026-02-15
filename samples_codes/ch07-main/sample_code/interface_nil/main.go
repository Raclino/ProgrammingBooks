package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface {
	Increment()
}

func main() {
	var pointerCounter *Counter
	fmt.Println(pointerCounter == nil) // prints true
	var incrementer Incrementer
	fmt.Println(incrementer == nil) // prints true
	incrementer = pointerCounter
	fmt.Println(incrementer == nil) // prints false
}
