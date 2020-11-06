package main

import "fmt"

// for 、if语句的变量 只能内部访问

func main2()  {
	// if条件语句
	if 1 > 2 {
		fmt.Println("1")
	}else if 1 > 3 {
		fmt.Println("2")
	}else {
		fmt.Println("0")
		return  // 跳出控制语句
	}

	// switch
	switch "lili" { 
		// switch func()+1 {  不一定是值 也可以是表达式  也可以是定义： switch name:= "lili"; {}
		// var x = interface{}
		// switch 还可以用于判断interface的变量类型  switch i := x.(type) {  


 		case "1":   // case的类型 必须和 switch类型一致
			// ...
		case "2":
			fallthrough // 穿透 继续执行下一个case
			// ...
		case "3":
			// ...
		case "4", "5", "6":  // 也可以是多个  但是为常量的时候 case不能重复
		default:
			// ...
		
	}

	// for 语句
	for i :=0; i< 10; i++ {
		// ...
		break // 跳出for循环
		continue // 跳出当前循环 继续执行下一个循环
	}

	// 遍历字符串和数组
	for i,val := range "hello" {
		// i, val

	}

	// while do while 实现 ；  go中没有while 和 do .. while
	for {
		if 1 > 0 {
			break   // while
		}
		// ... todo
	}

	for {
		// .... todo
		if 1 > 0 {
			break   // do ... while
		}
	}

}

