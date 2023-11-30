package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: function that returns a channel
// Channels are first-class values, just like strings or integers.

func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
	}
	fmt.Println("You're boring; I'm leaving.")
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
