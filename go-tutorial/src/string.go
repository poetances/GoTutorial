package main

import (
	"fmt"
	"strings"
)

// string虽然是值类型，但是内部是引用类型机制
func StringTutorial() {

	str1 := "this is str"
	fmt.Println(str1)

	str2 := `
	this is duo 
	this.is
	`
	fmt.Println(str2)

	// 1、长度
	strlen := len(str2)
	fmt.Println(strlen)

	// 2、拼接
	str3 := str1 + str2
	fmt.Println(str3)

	str4 := fmt.Sprintf("%v%v", str1, str2)
	fmt.Println(str4)

	// 3、strings包
	arr := strings.Split(str1, "-")
	fmt.Println(arr)

	joinstring := strings.Join(arr, "-")
	fmt.Println(joinstring)

	isContain := strings.Contains(str1, str2)
	fmt.Println(isContain)
}