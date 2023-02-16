package main

import (
	"fmt"
)

func foo() (int, string) {
	return 1, "OK"
}
func main() {
	//标准格式=变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号
	var a string
	a = "This is a string"
	fmt.Println(a)
	//批量声明
	var (
		b string
		c int
	)
	b = "This is a string"
	c = 10
	fmt.Println(b, c)
	//初始化
	var d int = 100
	fmt.Println(d)
	//多個同時初始化
	var str, isOK = "string", true
	fmt.Println(str, isOK)
	//類型推導
	var name = "Jack"
	fmt.Println(name)
	//在函數内部可以用變量
	n := 10
	fmt.Println(n)
	//匿名變量-如果你想省略某個值的時候可以用
	_, s := foo()
	fmt.Println(s)
	/*常量只是把var换成了const,但需要注意，
	如果声明多个常量时如果省略了值则表示和上面一行的值相同*/
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)
	/*iota为go语言的常量计数器只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。
	const中每新增一行常量声明将使iota计数一次*/
	const (
		num1 = iota
		num2
		num3
	)
	fmt.Println(num1, num2, num3)
	/*程序总输出:
	This is a string
	This is a string 10
	100
	string true
	Jack
	10
	OK
	100 100 100
	0 1 2*/

}
