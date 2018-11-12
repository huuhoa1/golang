package main

import (
	"fmt"
	
)
var (
	message string = "Hello Go!"
)

func main() {
	fmt.Println(message)
}

func init() {
	message = "Init"
}
