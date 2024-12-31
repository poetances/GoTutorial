package variable

import "fmt"

/*
int	整型，根据平台可能是32或64位
float	浮点型，float32和float64
bool	布尔型，true或false
complex	复数，complex64和complex128

byte	等价 uint8 可以表达ANSCII字符. 0-255
rune	等价 int32 可以表达Unicode字符. 0-0x10FFFF
string	字符串即字节序列，可以转换为[]byte类型即字节切片

衍生类型
数组（array）[5]int，长度为5的数组。长度是数组的一部分，所以[5]int和[10]int是不同的类型。
	数组是值类型，赋值和传参会复制整个数组。一般很少直接使用数组。
切片（slice）[]int，长度不固定的数组


常量（字面量）
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。
常量可以是字符、字符串、布尔值或数值。
*/

func Test() {
	var i int
	var f float64 = 12.0
	var b bool = true
	var c complex64 = 5 + 5i
	var bb byte = 255

	var arr [3]int = [3] int {1, 2}
	fmt.Println(i, f, b, c, bb, arr)
}

func Test2() {
	ii := 13
	var f float64 = 12

	var c  = 12
	const d = 12
	const e = 12.0
	const n = "name"
	const numExpression = 1 + 2
	
	// 批量声明
	var (
		age = 12
		name = "name"
	)

	const (
		age2 = 12
		name2 = "name"
	)

	// iota 常量重置，一般用于枚举
	const iota = 0
	const
	
	fmt.Println(ii, f, c, d, e, n, numExpression, age, name, age2, name2)

}