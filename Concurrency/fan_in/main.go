package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplexing
// These programs make Joe and Ann count in lockstep.
// We can instead use a fan-in function to let whosoever is ready talk.

func main() {
	c := fanInSelect(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // receive from c
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string) // create an unbuffered channel of type string
	go func() {            // launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // send data down the channel
			time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
		}
	}()
	return c // Return channel to the caller
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// Fan-in using select
// Rewrite our original fanIn function. Only one goroutine is needed. New:

func fanInSelect(input1, input2 <-chan string) <-chan string {
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
