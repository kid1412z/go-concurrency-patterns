package main

import (
	"fmt"
	"time"
)

// using quit chanel to receive dead message of a goroutine

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Second)
			case <-quit:
				fmt.Println("quit")
				quit <- "bye"
				return

			}
		}
	}()
	return c
}

func main() {
	quit := make(chan string)
	c := boring("Bob", quit)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	quit <- "quit"
	fmt.Println("Bob says:", <-quit)

}
