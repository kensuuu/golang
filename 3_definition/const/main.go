package main

import "fmt"

const Pi = 3.14
const (
	Usename  = "test_user"
	Password = "test_pass"
)
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Usename, Password)
	fmt.Println(big - 1)
}
