package main

import "fmt"

func main() {
	//双重循环
	//for i := 1; i <= 5; i++ {
	//	for j := 2; j <= 4; j++ {
	//		if i == 2 && j == 2 {
	//			continue //结束离它近的那个循环，继续离它的近的那个循环
	//		}
	//		fmt.Printf("i:%v,j:%v \n", i, j)
	//	}
	//}

	//加标签
lable2:
	for i := 1; i <= 5; i++ {
		//lable1: 如果标签没有使用，那么不用加，否则报错：定义未使用
		for j := 2; j <= 4; j++ {
			if i == 2 && j == 2 {
				continue lable2 //结束离它近的那个循环，继续离它的近的那个循环
			}
			fmt.Printf("i:%v,j:%v \n", i, j)
		}
	}
	fmt.Println("------ok")
}
