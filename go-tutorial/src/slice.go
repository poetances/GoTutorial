package main

import "fmt"

/*
var identifier []type

一个 slice 实际上是一个 结构体，包含：
  - 指针（Pointer）：指向底层数组的某个索引位置
  - 长度（Length）：slice 当前包含的元素个数
  - 容量（Capacity）：从 slice 开始位置到底层数组末尾的元素个数

	type SliceHeader struct {
	    Data uintptr // 指向底层数组的指针
	    Len  int     // 长度
	    Cap  int     // 容量
	}
但实际上，slice是引用类型，它的零值是 nil。

在 Go 中，make 是一个用于创建 slice（切片）、map（映射）和 channel（通道） 的内建函数。它用于初始化这些引用类型的数据结构，并分配底层存储。
*/
func SliceTutorial() {
	// init
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice1:", slice1)

	slice2 := make([]int, 5)
	fmt.Println("slice2:", slice2)

	var slice3 = []int{1, 2, 3, 4, 5}
	fmt.Println("slice3:", slice3)

	// append
	slice1 = append(slice1, 6)
	fmt.Println("slice1:", slice1)

	// copy
	slice4 := make([]int, 5)
	copy(slice4, slice2)
	fmt.Println("slice4:", slice4)

	// 截取
	subSlice1 := slice1[1:3]
	fmt.Println(subSlice1[:3]) // [1 2 3] (相当于 s[0:3])
	fmt.Println(subSlice1[2:]) // [3 4 5] (相当于 s[2:len(s)])
	fmt.Println(subSlice1[:])  // [1 2 3 4 5] (相当于 s[0:len(s)])
	fmt.Println("subSlice1:", subSlice1)


}
