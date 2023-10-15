package main

import "fmt"

func main() {
	//实现功能：如果口罩的库存小于30个，提示库存不足，否则提示：库存充足
	//定义口罩的数量
	var count int = 70

	if count < 30 {
		fmt.Println("口罩存量不足")
	} else {
		fmt.Println("库存充足")
	}
}
