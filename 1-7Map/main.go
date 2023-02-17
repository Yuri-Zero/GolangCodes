package main

import "fmt"

func main() {
	//map的使用
	student := make(map[string]int)
	student["小红"] = 1
	student["小李"] = 2
	fmt.Println(student)
	fmt.Println(student["小李"])
	//直接在声明时进行填充
	smap := map[string]int{
		"老韩": 3,
		"小伟": 4,
	}
	fmt.Println(smap)
	fmt.Println(smap["老韩"])
	//判断某键是否存在
	_, ok := smap["小伟"]
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	//使用delete函数删除键值对
	delete(smap, "小伟")
	fmt.Println(smap)
}
