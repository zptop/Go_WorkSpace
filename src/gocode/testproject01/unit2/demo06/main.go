package main

import "fmt"

func main() {
	//练习转义字符
	//\n换行
	fmt.Println("aaa\nbbb")

	//\b退格
	fmt.Println("aaa\bbbb")

	//\r光标回到本行的开头，后续输入就会替换原有的字符
	fmt.Println("aaaa\rbbb")

	//\t制表符
	fmt.Println("aaaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbb")
	fmt.Println("aaa\tbbbbb")

	//\" 原样输出
	fmt.Println("'golang'")
}
