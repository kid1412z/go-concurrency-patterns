package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel chain like this:
//
// input1 \
//         \
//          -- output
//         /
// input2 /

// so generating a random sequence

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
	go func() { for {c <- <-input1 } }()
	go func() { for {c <- <-input2 } }()
	return c
}

func main() {
	c := fanIn(boring("Alice"), boring("Bob"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
}
