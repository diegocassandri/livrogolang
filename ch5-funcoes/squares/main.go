// The squares program demonstrates a function value with state.
package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

}

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int { //Closure
		x++
		return x * x
	}
}
