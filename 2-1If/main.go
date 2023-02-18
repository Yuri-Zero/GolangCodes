package main

import "fmt"

func main() {
	x := false
	if x {
		fmt.Println("成立")
	} else {
		fmt.Println("不成立")
	}
	fmt.Println("执行完成")
	a := 100
	if a > 50 {
		if a > 100 {
			fmt.Println("a >100")
		} else {
			fmt.Println("a>50但小于100")
		}
	} else {
		fmt.Println("a<50")
	}
}
