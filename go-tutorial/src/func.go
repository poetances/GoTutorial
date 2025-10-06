package main

import "fmt"

func Maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func Swap(x, y string) (string, string) {
	return y, x
}

// 返回值命名
func Calc(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func Calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
 	return sum, sub
}

type CalcFunc func(int, int) int

// 函数作为参数
func calc3(x, y int, cb CalcFunc) int {
	return cb(x, y)
}

// 函数作为返回值。调用如下：calc4(1, 2)(4, 1)
func calc4(x, y int) CalcFunc {
	return func(i1, i2 int) int {
		return x + y
	}
}

/*
匿名函数定义：
func (参数) 返回值 { }

匿名函数的引入，主要是解决：在go中，普通函数是没法定义在函数内部的，不像swift有内部函数。
*/

func FuncTutorial() {
	var calc CalcFunc = Maximum
	fmt.Println(calc(1, 2))

	fmt.Printf("%T\n", Maximum)

	// 匿名函数
	a := calc3(12, 14, func(i1, i2 int) int {
		return i1 + i2
	})
	fmt.Println(a)
}
   