package main

import "fmt"

func main() {
	//定义一个字符串
	var str string = "hello golang你好"

	//方式1：普通for循环：按照字节进行遍历输出的（暂时先不使用中文）
	//for i := 0; i < len(str); i++ { //i可以理解为字符串的下标
	//	fmt.Printf("%c", str[i])
	//}

	//方式2：for range
	//对str进行遍历，遍历的每个结果的索引值被i接收，每个结果的具体数值被value接收
	for i, value := range str {
		fmt.Printf("索引为：%d,具体值为：%c \n", i, value)
	}
}
