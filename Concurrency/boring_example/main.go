package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go lessBoring("boring...")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

// Initial func
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

// Make the intervals between messages unpredictable (still under a second).
func lessBoring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
	}
}
