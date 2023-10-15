package main

import "fmt"

func main() {
	//可以通过指针改变指向值
	var num int = 10
	fmt.Println(num)

	var ptr *int = &num
	*ptr = 20
	fmt.Println(num)

	//指针变量接收的一定是地址值
	//var ptr2 *int = num  错误的地址，应该是&num
	//fmt.Println(ptr2)

	//指针变量的地址不可以不匹配
	var num2 int = 20
	fmt.Println(num2)
	//var ptr3 *float32 = &num2  *float32意味着这个指针指向的是float32类型的数据，但是&num2对应的是int类型的不可以。
	//fmt.Println(ptr3)

	//基本数据类型（又叫值类型），都有对应的指针类型，形式为 *数据类型。比如int对应的指针就是 *int，float32对应的指针类型就是 *float32,依次类推

}
