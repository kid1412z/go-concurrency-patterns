package main

import (
	"fmt"
	"math/rand"
	"time"
)

// using select to deal with output

// channel chain like this:
//
// input1 \
//         \
//          -- output
//         /
// input2 /

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

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Alice"), boring("Bob"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
}
