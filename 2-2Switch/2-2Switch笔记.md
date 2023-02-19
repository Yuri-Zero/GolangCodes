# Switch语句笔记
switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
但不同于C/C++ Java等支持Switch语法的编程语言，Go只执行符合条件的case，而非之后所有的case，这是因为Golang会在每个case后自动提供break语句。除非以 **fallthrough** 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
```go
	score := 70
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
```
输出结果:
```go
C
```
```go
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
```
输出结果:
```go
c
D
```
同时Golang的switch语法还支持一种无条件switch.这其实可以被视为是switch true的省略写法
```go
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```
等同于
```go
	t := time.Now()
	switch true {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```
## type switch
switch 语句还可以被用于 type-switch 来判断某个 **interface{}** 变量中实际存储的变量类型。
>interface{} 类型，空接口，是导致很多混淆的根源。interface{} 类型是没有方法的接口。由于没有 implements 关键字，所以所有类型都至少实现了 0 个方法，所以 所有类型都实现了空接口。这意味着，如果您编写一个函数以 interface{} 值作为参数，那么您可以为该函数提供任何值。
```go
switch x.(type){
    case type:
       statement(s)      
    case type:
       statement(s)
    default: 
       statement(s)
}
```
例:
```go
func foo(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	}
}
func main(){
    var x float32
    var y int
	foo(x)
    foo(y)
}
```
结果为:
```go
float32
int
```