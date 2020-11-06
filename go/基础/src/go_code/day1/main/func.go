package main

import (
	nameUtil "../utils" // 可以重新命名
	"fmt"
	"strconv"
)

// 12. 全局匿名函数
var (
	Func1 = func (a int) int {
		return a
	}
)


// 11. init函数
func init() {
	fmt.Println("初始化执行，在main函数之前调用") // 执行顺序 global Var -> init() -> main()
}


func mains()  {
	// 1.
	var result string = f1("lii", 30)
	fmt.Println(result) // lii30
	// 2.
	isCallback := nameUtil.Fors() // 调用公共方法

	// 3.
	if isCallback {
		// main() // 递归调用
		weight := "20kg"
		f2(&weight)
		// go默认是 拷贝值不是地址 ，不会改变原理的变量，如果想改可以传地址
		fmt.Printf("地址值=== %s", weight)
	}

	// 5. 函数可以赋给变量
	ff := f1
	ff("xiaoming", 30)

	// 6. 也可以作为参数传递
	nums := funcVar(add, 10)
	fmt.Printf("\n结果 %d", nums)

	// 9.
	ignore()
	// 10.
	argsm(10, 20, 30, 40)

	// 12 匿名函数
	// -1. 匿名函数返回值调用
	funcni := func (a int) int {
		return a
	}(10)
	fmt.Println(funcni) // 10

	// -2. 函数表达式
	funcnif := func(b int) int {
		return b
	}
	ps := funcnif(10)
	fmt.Println(ps) // 10

	// 全局匿名函数
	Func1(100)

	//13. 闭包
	bibaos := bibao()
	fmt.Println(bibaos(10)) // 11
	fmt.Println(bibaos(11)) // 22

	//14. defers
	defers()
}





// 1.
func f1(name string, age int64) string {
	return name + strconv.FormatInt(age, 10) // 告知进制
}

// 4.
// func f1() { } // go 函数不支持重载

//

// 3.传递地址时，操作后原来的值也收到影响
func f2(weight *string) int64 {
	var weg_ int64
	if weg, err := strconv.ParseInt(*weight, 10, 64); err != nil {
		weg_ = weg
		*weight = "1"
		fmt.Printf("%T, %v\n", weg, *weight)
	}
	return weg_
}


// 6. 函数作为参数传递
func funcVar(funcvar func(int) int, num int) int {
	return funcvar(num)
}

// 7.支持自定义类型
type myFuncType func(int) int
func funcVarSelfType(funcvar myFuncType, num int) int {
	return funcvar(num)
}

func add(num int) int {
	return num
}

// 8 支持对函数返回值命名
func returnName(ages int) (name string, age int) {
	return name, age
}


// 9. _忽略返回值
func ignore() {
	age, _ := calc(100)  // 忽略age的fan
	fmt.Printf("\n返回结果 %d", age)

}

func calc(num int) (age int, weight int) {
	age = num +1
	weight = num + 2
	return
}

// 10. 支持参数列表中有可变参数
func argsm(n int, args... int) {
	fmt.Println(n, args) // 10 [20 30 40]
	fmt.Println(args[1]) // 30
}

// 13. 闭包
func bibao() func (int) int {
	var a int = 1
	return func (b int) int {
		a+=b
		return a
	}
}


// 14.defer延迟机制 ： 执行到defer语句时 不执行，而是压入栈中，等函数执行完毕后再执行defer语句，安装先入后出的原则, 如果有值，则拷贝值
func defers() {
	defer fmt.Println(2) // 后执行
	defer fmt.Println(1) // 先执行

	de := 10
	fmt.Println(de)
}

















