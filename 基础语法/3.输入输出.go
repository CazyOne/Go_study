package main

import "fmt"

func main() {
	// 输入输出函数
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)            // Scanln函数用于从标准输入读取一行数据，并将其存储到变量中
	fmt.Printf("你好，%s！\n", name) // Printf函数用于格式化输出
}
