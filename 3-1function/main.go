package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

type subTool func(x int, y int) int

func sub(fn subTool, x int, y int) int {
	return fn(x, y)
}
func AddUpper() func(x int) int {
	n := 10
	return func(x int) int {
		n = n + x
		return n
	}
}
func main() {
	x := 10
	y := 10
	z := add(x, y)
	fmt.Println(z)
	sub(func(x int, y int) int { return x - y }, x, y)
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
}
