// new() and make()
// GC happends automatically

package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println("val:", ptr)
	// default value of pointer is nil

	num := 10
	var ptr = &num

	fmt.Println("val :", ptr)
	fmt.Println("val:", *ptr)

	// *ptr -> actual value of pointer
	// ptr -> address of pointer

	*ptr = *ptr + 2
	fmt.Println("new val:", num)
	// always changes the original value rather than copy of it.
}
