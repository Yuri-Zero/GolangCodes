package main

import "fmt"

func main() {
	//一.切片创建
	//1.切片声明
	var s1 []int //<-全局声明
	_ = s1
	//2.:=短变量声明
	s2 := []int{} //<-局部短变量声明
	_ = s2
	//3.make
	/*make(Type,len,cap)
	第一种，只传类型，不指定实际占用的内存空间和提前预留的内存空间，适用于 map
	 和 channel 。
	第二种，指定实际占用的内存空间为 2，不指定提前预留的内存空间。
	第三种，指定实际占用的内存空间为 2，指定提前预留的内存空间是 4。*/
	var s3 []int = make([]int, 0)
	fmt.Println(s3)
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	//4.初始化赋值
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	//5.数组切片
	arr := [4]int{1, 2, 3, 4}
	s6 := arr[1:3]
	fmt.Println(s6)
	//切片初始化
	var arr1 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//全局初始化
	var slice1 []int = arr1[0:10] //自arr[0,10)切片
	fmt.Println(slice1)
	var slice2 []int = arr1[:10] //自arr[0,10)切片
	fmt.Println(slice2)
	var slice3 []int = arr1[0:] //自arr[0,len(arr1))切片
	fmt.Println(slice3)
	var slice4 []int = arr1[:] //自arr[0,len(arr1))切片
	fmt.Println(slice4)
	var slice5 []int = arr1[:len(arr1)] //自arr[0,len(arr1))切片
	fmt.Println(slice5)
	//局部初始化就是换成短变量声明
	//使用make函数进行初始化
	slice6 := make([]int, 10, 10) //创建一个长度为10，容量为是10，元素为0的切片
	fmt.Println(slice6)
	arr2 := [3]int{1, 2, 3}
	slice7 := arr2[2:3]
	fmt.Println(slice7) //该切片元素为[3]
	fmt.Println(arr2)   //该数组元素为[1,2,3]
	slice7[0] += 100
	fmt.Println(slice7) //该切片元素变为[1,103]
	fmt.Println(arr2)   //该数组元素变为[1,2,103]
	slice7 = append(slice7, 104)
	fmt.Println(slice7) //该切片变为[103,104]
	fmt.Println(arr2)   //底层数组元素不变
	slicecopy1 := []int{1, 2, 3}
	slicecopy2 := make([]int, 10)
	fmt.Println("slicecopy1:", slicecopy1)
	fmt.Println("slicecopy2:", slicecopy2)
	//切片拷贝使用copy函数，长度由小的决定，可指向同一底层数组，允许元素区间重叠
	//slicecopy1: [1 2 3]
	//slicecopy2: [0 0 0 0 0 0 0 0 0 0]
	copy(slicecopy2, slicecopy1)
	fmt.Println("copy操作执行后的slicecopy1", slicecopy1)
	fmt.Println("copy操作执行后的slicecopy2", slicecopy2)
	//copy操作执行后的slicecopy1 [1 2 3]
	//copy操作执行后的slicecopy2 [1 2 3 0 0 0 0 0 0 0]
}
