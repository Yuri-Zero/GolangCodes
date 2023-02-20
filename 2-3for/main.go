package main

import "fmt"

func main() {
	// sum := 0
	// i := 0
	// for i < 10 {
	// 	sum += i
	// 	i++
	// }
	// fmt.Println(sum)
	for i := 0; i <= 5; i++ {
		fmt.Println("i:", i)
		for j := 0; j <= 6; j++ {
			fmt.Println("j:", j)
			continue
		}
	}
	var a int = 10
	for a < 25 {
		if a == 20 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
always:
	fmt.Println("一直会执行这一行")
	goto always
	fmt.Println("永远不会执行到这里")
}
