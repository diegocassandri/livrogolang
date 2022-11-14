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
	canal := make(chan string)

	//T2
	go func() {
		canal <- "Veio da t2"
	}()

	// T1
	msg := <-canal
	fmt.Println(msg)

}
