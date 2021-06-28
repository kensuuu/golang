package main

import "fmt"

//varは関数の外にも宣言できる
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

// 型を宣言したい場合はvarを使う
func foo() {
	xi := 1
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}
func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
