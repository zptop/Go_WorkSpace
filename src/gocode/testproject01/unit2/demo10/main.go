package main

import "fmt"

func main() {
	//进行类型转换
	var n1 int = 100
	//var n2 float32 = n1 在这里自动转换不好使，比如显式转换
	fmt.Println(n1)
	//fmt.Println(n2)
	var n2 float32 = float32(n1)
	fmt.Println(n2)

	//注意：n1的类型其实还是int类型，只是将n1的值100转为了float32而已，n1还是int的类型
	fmt.Printf("%T", n1) //int
	fmt.Println()

	//将int64转为int8的时候，编译不会出错的，但是数据会溢出
	var n3 int64 = 888888
	var n4 int8 = int8(n3)
	fmt.Println(n4)

	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30 //一定要匹配 “等号” 左右的数据类型
	fmt.Println(n5)
	fmt.Println(n6)

	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127 //编译通过，但是结果可能会溢出
	//var n9 int8 = int8(n7) + 128 //编译不会通过

	fmt.Println("n8:", n8)
	//fmt.Println("n9:", n9)
}
