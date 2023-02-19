package main

import (
	"fmt"
	"time"
)

func foo(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	}
}
func main() {
	score := 70
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
		fallthrough
	default:
		fmt.Println("D")
	}
	t := time.Now()
	switch true {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	var x float32
	var y int
	foo(x)
	foo(y)
}
