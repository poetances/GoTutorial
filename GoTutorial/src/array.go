package main

import "fmt"

/*
数组和切片非常类似。但是数组的长度是固定的，切片的长度是可变的。

var arrayName [size]dataType
*/
func ArrayTutorial() {
	var arr1 []int
	fmt.Println("emp:", arr1)

	var arr2 = []int{1, 2, 3, 4, 5}
	fmt.Println("arr2:", arr2)
}