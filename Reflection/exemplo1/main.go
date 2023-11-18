package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)  // um reflect.Type
	fmt.Println(t.String()) // int
	fmt.Println(t.Kind())   // int
	fmt.Println(t)          // int

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File

	v := reflect.ValueOf(3) // um reflect.Value
	fmt.Println(v)          // 3
	fmt.Printf("%v\n", v)   // 3
	fmt.Println(v.String()) // <int Value>

	j := v.Type()           // um reflect.Value
	fmt.Println(j.String()) //int
}
