package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello world"[0]))

	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))
}