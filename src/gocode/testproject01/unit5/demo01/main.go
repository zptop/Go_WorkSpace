package main

import "fmt"

//func 函数名(形参列表)(返回值类型列表){
//执行语句
//return + 返回值列表
//}

// 自定义函数：功能：两个数相加
func cal(num1 int, num2 int) int { //如果返回值类型就一个，那么()是可以省略不写的
	var sum int = 0
	sum = num1 + num2
	return sum
}
func main() {
	var sum = cal(10, 20)
	fmt.Println(sum)
}
