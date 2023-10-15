package main

import "fmt"

func main() {
	//定义浮点类型的数据
	var num1 float32 = 3.14
	fmt.Println("num1:", num1)

	//可以表示正浮点数，也可表示负浮点数
	var num2 float32 = -3.14
	fmt.Println("num2:", num2)

	//浮点数可以用十进制表示，也可以用科学计数法表示，E不区分大小写
	var num3 float32 = 314e-2
	fmt.Println("num3:", num3)

	var num4 float32 = 314e+2
	fmt.Println("num4:", num4)

	//浮点数可能会有精度的损失，所以通常情况下，建议使用float64
	var num7 float32 = 256.000000916
	fmt.Println("num7:", num7)
	var num8 float64 = 256.000000916
	fmt.Println("num8:", num8)

	//golang中默认的浮点类型为float64
	var num9 = 3.14
	fmt.Printf("num9对应的默认类型为：%T", num9)
}
