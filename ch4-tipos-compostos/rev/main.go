package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Revertendo um Array manualmente
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//Revertendo uma fatia (slice)
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2]) //Passa apenas os dois primeiros indices
	reverse(s[2:]) //passa todos a partir do segundo indice
	reverse(s)     // Passa a fatia inteira
	fmt.Println(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	// Interactive teste os reverse
	input := bufio.NewScanner(os.Stdin)

outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
