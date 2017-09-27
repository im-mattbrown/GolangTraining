package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//function to fetch the body of an html page and print its resulting html
func main() {
	res, _ := http.Get("http://mattericbrown.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
