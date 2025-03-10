package main

import "fmt"

/*
defer语句很好理解，跟swift中的defer一样，一般用于资源回收的。
*/
func DeferTutorial() {
	fmt.Println("开始")
	defer fmt.Println(1)
	fmt.Println(2)
	fmt.Println("结束")
}