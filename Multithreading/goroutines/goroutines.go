package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}

// T1 -> main
func main() {
	go task("A") // go routines green threads
	go task("B")
	task("C")
}
