package main

import "fmt"


func main() {
	testArray()
}

/*
数组和切片非常类似。但是数组的长度是固定的，切片的长度是可变的。

*/
func testArray() {
	var arr1 [5]int
	fmt.Println("emp:", arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2:", arr2)
}