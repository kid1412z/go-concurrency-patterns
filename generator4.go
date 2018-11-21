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

// there is a wait channel used to synchronize
// so the output is fixed sequence

type Message struct {
	str string
	wait chan bool
}

// function returns a receive-only channel of strings
func boring(msg string) <-chan Message {
	c := make(chan Message)
	wait:=make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			wait <- true // block
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() { for {c <- <-input1 } }()
	go func() { for {c <- <-input2 } }()
	return c
}

func main() {
	c := fanIn(boring("Alice"), boring("Bob"))
	for i := 0; i < 20; i++ {
		msg1:=<-c
		fmt.Println(msg1.str)
		<-msg1.wait

		msg2:=<-c
		fmt.Println(msg2.str)
		<-msg2.wait
	}
}
