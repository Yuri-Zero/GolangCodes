package main

import "fmt"

type Student struct {
	id   int
	name string
}
type Astruct struct {
	id   int
	name string
}
type Bstruct struct {
	id int
}
type Cstruct struct {
	Astruct
	Bstruct
}

func main() {
	// 自定类型
	type Myint int
	//类型别名
	type Aint = int
	//检验区别
	var a Myint
	var b Aint
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of a:%T\n", b)
	//声明结构体变量
	var xiaoMing Student
	//访问/修改字段值
	xiaoMing.id = 5050
	xiaoMing.name = "小明"
	fmt.Println(xiaoMing)
	ptr := &xiaoMing.id
	*ptr = 6060
	fmt.Println(xiaoMing)
	var ptrStudentStruct *Student = &xiaoMing
	ptrStudentStruct.name = "小李"
	fmt.Println(xiaoMing)
	var stu1 Student = Student{12, "小孙"}
	fmt.Println(stu1)
	var stu2 Student = Student{name: "小张"}
	fmt.Println(stu2)
	var c Cstruct
	c.Astruct.id = 100 //valid
	c.Bstruct.id = 100 //valid
	//c.id = 100 <- invalid
	c.name = "C" //<- OK

}
