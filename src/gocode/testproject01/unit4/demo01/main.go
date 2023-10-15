package main

import "fmt"

func main() {
	//实现功能：如果口罩的库存小于30个，提示库存不足
	//var count int = 20
	//单分支
	//if count < 30 {
	//	fmt.Println("口罩存量不足")
	//}

	//if后面表达式，返回结果一定是true或false
	//如果返回结果为true的话，那么｛｝中的代码就会执行
	//如果返回结果为false的话，那么｛｝中的代码就不会执行
	//if后面一定要有空格，和条件表达式分隔开
	//｛｝一定不能省略
	//条件表达式左右的（）是建议省略的
	//在golang里，if后面可以并列的加入变量的定义

	if count := 20; count < 30 {
		fmt.Println("口罩存量不足")
	}
}
