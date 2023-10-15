package main

import "fmt"

func main() {
	//双重循环
lable2:
	for i := 1; i <= 5; i++ {
		//lable1: 如果标签没有使用，那么不用加，否则报错：定义未使用
		for j := 2; j <= 4; j++ {
			fmt.Printf("i:%v,j:%v \n", i, j)
			if i == 2 && j == 2 {
				break lable2 //结束指定标签对应的循环
			}
		}
	}
	fmt.Println("------ok")
}
