package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Synchronization
// When the main function executes <–c, it will wait for a value to be sent.

// Similarly, when the boring function executes c <– value, it waits for a receiver to be ready.

// A sender and receiver must both be ready to play their part in the communication. Otherwise we wait until they are.

// Thus channels both communicate and synchronize.

// Note for experts: Go channels can also be created with a buffer.

// Buffering removes synchronization.

func main() {
	c := make(chan string)
	go boring("message", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring; I'm leaving.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
	}
}
