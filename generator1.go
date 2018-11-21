package main

import (
	"fmt"
	"math/rand"
	"time"
)

// start goroutine in a function generating data to a channel
// return the receive-only channel to caller

// function returns a receive-only channel of strings
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("boring")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
