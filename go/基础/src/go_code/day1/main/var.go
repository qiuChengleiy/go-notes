package main

import (
	//"fmt"

//	"../test"
)

// 变量声明： ========================================================================

var a1, a2, a3 int
var ai = 1
var (
	a4 = 5
	a5 = 6
)



// func vars() {
// 	// 变量生命
// 	// 1. 声明类型不赋值  使用默认值 
// 	var i int // int 默认是0
// 	// 2. 类型自动转换
// 	var num = 123 
// 	// 3. := 声明 
// 	name := "lili" // var name = "lili" 并且 name 没有被声明过 // string 默认是''
// 	// 4. 多变量声明
// 	var a, b, c int
// 	var d, e, f = 1, 2, 3 // 可以使用类型推导和赋值
// 	h, i, j := 4, 5, 6
// }



/* 数据类型： ========================================================================

基本数据类型： 
			 数值类型 ： int 整数型   float32,float 64 浮点型 default: 0
			 字符型 
			 字符串 string default: ''
			 布尔类型 bool default: false

引用类型/ 派生/ 复杂数据类型：
			指针：pointer
			数组
			结构体 struct
			管道   channel
			函数	func
			切片	slice	
			接口	Interface
			map

*/

// func dataVar() {
// 	// 类型判断
// 	fmt.Println("typeof %T", 100) // %T 返回类型
// 	// import "unsafe"
// 	// unsafe.Sizeof(varname) 返回变量占用的字节数

// 	// 浮点型 float64 精确度要高
// 	var price float64 = 12.00001

// 	// string类型不可修改
// 	var str = "hello"
// 	// str[0] = "j" 报错
// 	// "\n" 会识别转义字符 ``会以原生形式返回 包括换行符
// 	demo := `
// 		var a = 1
// 	`
// 	// 字符串拼接
// 	name := "he" + "lo" + // 这个+ 必须保留在上边 如果拼接太多的时候
// 	"wor" + "d"

// 	// 数据类型转换: golang 不能自动转换 需要显示
// 	// 基本类型转字符串类型
// 	// func Sprintf
// 	// exa: 
// 	boos := false
// 	nums := 1
// 	floats := 1.1
// 	strs := "hello"
// 	var type_1 = fmt.Sprintf("%t", boos)
// 	var type2 = fmt.Sprintf("%d", nums)
// 	var type_3 = fmt.Sprintf("%f", floats)
// 	var type_4 = fmt.Sprintf("%c", strs)

// 	// 方法2： strconv
// 	// import strconv strconv.format
// 	// var type_1 = strconv.FormatBool(boos)
// 	// var type2 = strconv.FormatInt(nums)
// 	// var type_3 = strconv.FormatFloat(floats)

// 	// string 转 基本数据类型
// 	// strconv.ParseBool
// 	// strconv.ParseInt
// 	// strconv.ParseFloat

// 	// 注意：字符串转的值 "123" -> 123  "helllo" -> int:0 float:0 bool: false 





// 	// 指针类型
// 	var ptr int = 3 // &ptr 这个就是内存地址
// 	var pt *int = &ptr // *int 就是指针类型 pt -> ptr
// 	fmt.Println("pt指向的值是 %v", *pt) // *pt就是指向的值 注意数据类型

// 	// test.go
// 	// package test
// 	// var Test = "test"   变量名称大写可以被导出使用 首字母大写公有， 首字母小写私有  golang没有public private关键字

// 	// fmt.Println(test.Test)

// }