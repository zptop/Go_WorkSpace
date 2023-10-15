package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义一个整数类型
	var num1 int8 = 23
	fmt.Println(num1)

	var num2 uint8 = 20
	fmt.Println(num2)

	var num3 = 28
	//Printf的作用是：格式化的，把num3的类型填充的 %T 的位置上
	fmt.Printf("num3的类型是:%T", num3)
	fmt.Println()
	//占用的字节数
	fmt.Println(unsafe.Sizeof(num3))
}
