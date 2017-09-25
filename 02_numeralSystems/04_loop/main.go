package main

import "fmt"

func main() {
	// basic for loop printing numbers 1-199 in decimal binary and hex
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
