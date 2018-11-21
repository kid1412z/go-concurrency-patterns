package main

import (
	"fmt"
)

// channel used to communicate and synchronize between goroutines
func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println("value=", <-ch)
}
