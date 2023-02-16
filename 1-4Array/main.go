package main

import "fmt"

//全局初始化
//在初始化时右侧指定长度
var arr = [5]int{1, 2, 3, 4, 5}

//自动推导长度
var arr2 = [...]int{1, 2, 3}

//指定索引内容
var arr3 = [4]int{2: 100, 3: 99}

//全局局部数组
var arr4 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

//值传递数组
func foo1(arr [5]int) {
	arr[0] = 10
	fmt.Println(arr) //函数内形参的数值会被修改，原本实参的不会，因为是值传递
}

//指针传递数组
func foo2(arr *[5]int) {
	arr[0] = 10
}
func main() {
	//局部初始化
	//在初始化时右侧指定长度
	b := [5]int{1, 2, 3, 4, 5}
	//自动推导长度
	c := [...]int{1, 2, 3}
	//指定索引内容
	d := [4]int{2: 100, 3: 99}
	fmt.Println(arr, arr2, arr3)
	fmt.Println(b, c, d)
	//局部二维数组初始化
	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr4)
	fmt.Println(e)
	//foo1中参数为值传递(golang默认数组为值传递)，修改只会影响拷贝，不会影响到本身
	foo1(arr)
	fmt.Println(arr)
	//foo2为指针传递，会影响到本身
	foo2(&arr)
	fmt.Println(arr)
}

/*
代码全部输出:
[1 2 3 4 5] [1 2 3] [0 0 100 99]
[1 2 3 4 5] [1 2 3] [0 0 100 99]
[[1 2 3] [7 8 9]]
[[1 2 3] [4 5 6]]
[10 2 3 4 5]
[1 2 3 4 5]
[10 2 3 4 5]
*/
