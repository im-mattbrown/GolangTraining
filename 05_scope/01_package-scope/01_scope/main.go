//package level scope, encloses the other scopes main() and foo()
package main

import "fmt"

// declared at package level, available to inner level scope
var x int = 42

//inner scope from { to } x is available
func main() {
	fmt.Println(x)
	foo()
}

// inner scope where x is available
func foo() {
	fmt.Println(x)
}
