package main

import "fmt"

func main() {
	//1.定义一个字符串
	var s1 string = "全面拥抱golang"
	fmt.Println(s1)

	//2.字符串是不可变的，指的是字符串一旦定义好，其中的字符的值不能改变
	var s2 string = "abc"
	//s2[0] = "t"
	fmt.Println(s2)

	//3.字符串表示形式
	//(1)如果字符串中没有特殊字符，字符串的表示形式用双引号
	//(2)如果字符串中有特殊字符，字符串的表示形式用反引号
	var s4 string = `
		package main
		
		import "fmt"
		
		func main() {
			//测试布尔类型的数值
			var flag01 bool = true
			fmt.Println(flag01)
		
			var flag02 bool = false
			fmt.Println(flag02)
		
			var flag03 bool = 5 < 9
			fmt.Println(flag03)
		}`
	fmt.Println(s4)
}
