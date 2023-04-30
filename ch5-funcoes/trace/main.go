package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	//... Lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
