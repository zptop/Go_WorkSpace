package main

import "fmt"

var n7 = 100
var n8 = 9.7

// 设计者认为上面的全局变量的写法太麻烦了，可以一次性声明
var (
	n9  = 50
	n10 = "netty"
)

func main() {
	//第一种
	var num int = 18
	fmt.Println(num)

	//第二种：指定变量的类型，但是不赋值，使用默认值
	var num2 int
	fmt.Println(num2)

	//第三种：如果没有写变量的类型，那么会根据=后面的值进行判定变量的类型（自动类型推断）
	var num3 = 10
	fmt.Println(num3)

	//第四种：省略var，注意 := 不能写为 =
	sex := "男"
	fmt.Println(sex)

	fmt.Println("---------------------------------------------------------------------")
	//声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	n6, height := 6.9, 100.6
	fmt.Println(n6)
	fmt.Println(height)
	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)
	fmt.Println(n10)
}
