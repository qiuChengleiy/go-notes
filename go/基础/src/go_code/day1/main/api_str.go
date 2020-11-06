package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1. 字符串常用函数
	strs()
}

func strs() {
	// 统计字符串长度
	fmt.Println(len("hello")) // 5

	// 遍历字符串
	strs := "hello你好"
	r := []rune(strs) // [5/32]0xc00010eec0104 同时处理中文问题
	print(r)
	for i :=0; i < len(r); i++ {
		fmt.Printf("\n字符=%c", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Printf("转换的错误是", err)
	}else {
		fmt.Printf("转换的结果是 %d", n)
	}

	// 整数转字符串
	strsm := strconv.Itoa(1111)
	fmt.Printf("\n %s", strsm) // "111"

	// 字符串转byte
	bytss := []byte("hellp")
	fmt.Println(bytss) // [104 101 108 108 112]

	// byte 转字符串
	strby := string(bytss)
	fmt.Println(strby) // "hellp"

	// 10进制转2，8，16
	jinzhi := strconv.FormatInt(100, 2)
	fmt.Println(jinzhi) // 1100100

	// 字符串查找
	strSearch := strings.Contains("hello", "lo")
	fmt.Println(strSearch) // true

	// 统计字符串指定字符的个数
	strCount := strings.Count("hello", "l")
	fmt.Println(strCount) // 2

	// 不区分大小写比较
	bijiaoAa := strings.EqualFold("ABC", "aBc")
	fmt.Println(bijiaoAa) // true

	// 字符串大小写转换
	strUp := strings.ToUpper("hello")
	fmt.Println(strUp) // HELLO
	strLow := strings.ToLower(strUp)
	fmt.Println(strLow) // hello

	// 返回查找索引
	indexs := strings.Index("hello", "h") // 第一次出现
	fmt.Println(indexs) // 0
	indexsLast := strings.LastIndex("hello", "l") // 最后一次出现
	fmt.Println(indexsLast) // 3

	// 字符串替换
	strReplace := strings.Replace("hello", "l", "L", -1) // n表示替换几个字符，  -1表示全部替换
	fmt.Println(strReplace) // heLLo

	// 字符串分割成数组
	arrStrSplit := strings.Split("hello world", " ")
	fmt.Println(arrStrSplit) // [hello world]

	// 去除两端空格
	strTrim := strings.TrimSpace(" hello world ")
	fmt.Println(strTrim) // hello world

	// 去除两端指定字符串
	strStrTrim := strings.Trim("! hello world !", "!")

	// 去除左端指定字符
	strStrLTrim := strings.TrimLeft("! hello world !", "!")
	// 去除右端指定字符
	strStrRTrim := strings.TrimRight("! hello world !", "!")

	// 判断字符串是否以给定字符开头
	strIsIndex := strings.HasPrefix("http://", "http")
	// 判断字符串是否以给定字符串结尾
	strIsIndex := strings.HasSuffix("NTL,SSS", "S")
}







