package main

import (
	"fmt"
	"math/rand"
	"time"
)

// using select to implement timeout of each message

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := boring("Bob")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(700 * time.Millisecond): // timeout each massage
			fmt.Println("timeout")
			return
		}
	}
}
