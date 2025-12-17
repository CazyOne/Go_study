package main

import "fmt"

func start() {
	fmt.Println(globalVar)
}

var globalVar string = "变量的定义"

func main() {
	// 变量的定义格式：var 变量名 变量类型 = 变量值 3种方式
	var name string = "张三"
	var age = 18
	isStudent := true
	start()
	fmt.Println("姓名：", name)
	fmt.Println("年龄：", age)
	fmt.Println("是否是学生：", isStudent)
}
