package main

import (
	"fmt"
	"math/rand"
	"time"
)

// start a goroutine does not make caller wait

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// does not make caller wait
	go boring("boring")
}
