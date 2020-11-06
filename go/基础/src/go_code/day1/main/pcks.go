package main

import "fmt"

func main1() {
	// fmt.ScanIn() 监听用户输入
	var name string
	var age int
	// fmt.Println("请输入name")
	// fmt.Scanln(&name)
	// fmt.Printf("名字是 %v", name) // Printf 格式化输出

	// Scanf 指定格式输入
	fmt.Println("请输入name age")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("名字是 %v \n 年龄是 %v", name, age)
	// 请输入name age
	// lili 20
	// 名字是 lili 
	// 年龄是 20% 
}