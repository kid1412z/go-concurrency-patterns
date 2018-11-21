package main

import (
	"fmt"
)

// put data to channel will block current goroutine
func main() {
	ch := make(chan int)
	// block and dead
	ch <- 1
	fmt.Println("value=", <-ch)
}
