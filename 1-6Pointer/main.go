package main

import "fmt"

func foo1(x int) {
	x += 100
}
func foo2(x *int) {
	*x += 100
}
func main() {
	var (
		num int  = 10
		ptr *int = &num
	)
	fmt.Printf("变量num的地址为:%x\t 指针变量ptr指向的地址为:%x\t指针变量ptr的地址为%x", &num, ptr, &ptr)
	fmt.Println("变量num的值:", num)
	fmt.Println("指针取值:", *ptr)
	num2 := *ptr //注，此时ptr指向num的地址
	fmt.Println("num2:", num2)
	num3 := 10
	//值传递
	foo1(num3)
	fmt.Println(num3) //实参的值并不会改变
	//指针传递
	foo2(&num3)
	fmt.Println(num3) //实参的值变为110
	//内存分配-new函数
	a := new(int)
	fmt.Println(a)
}
