package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		// %q for UTF-8
		fmt.Printf("%q \n", i)
	}
}
