package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "hello"
	i, _ := strconv.Atoi(s)
	// if err != nil{
	// 	fmt.Println("ERROR")
	// }
	fmt.Printf("%T %v\n", i, i)
}