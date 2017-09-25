package main

import "fmt"

func main() {
	//%b for binary representation %x for hexadecimal representation
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//%#x to print 0x before hex value
	fmt.Printf("%d - %b - %#x \n", 33, 33, 33)
	//capital X for 0X before hex value
	fmt.Printf("%d - %b - %#X \n", 44, 44, 44)
	// \t for tabs in between
	fmt.Printf("%d \t %b \t %#X \n", 69, 69, 69)
}
