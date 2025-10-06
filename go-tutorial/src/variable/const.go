package variable

import (
	"fmt"
)

func ConstTutorial() {

	const pi float32 = 3.1415926
	fmt.Println("常量声明", pi)

	const (
		a = "A"
		b = 1
	)
	fmt.Println(a, b)

	// ​常量组的工作机制​：在同一个 const声明块中，如果一行常量省略了值，它会自动继承并复制前一个常量的表达式和计算结果​
	const (
		c = 1
		d
		e
	)
	fmt.Println(c, d, e) // 输出结果， 1， 1， 1

	/*
		var (
			f = 1
			i
			j
		)
		fmt.Println(f, i, j)
		直接报错，在 var块中声明变量时，​每一个变量都必须有明确的初始化值或明确的类型​

		var (
			f = 1
			i int
			j string
		)
		这样是可以的
	*/
}

func ConstTutorial2() {
	// iote const一般组合使用
	// const iota = 0 // Untyped int.
	// 从定义中，我们也可以看出。iota默认就是0
	const i = iota
	fmt.Println(i)

	// const 的基本规则：
	// 基本递增​：const块中每增加一行，iota值加 1（从 0 开始）
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d) // 0 1 2 3

	// 跳过某些值​：使用 _（空白标识符）可以跳过特定的 iota值。
	const (
		e = iota
		f
		j
		_ // 3被特定跳过
		h
	)
	fmt.Println(e, f, j, h) // 0 1 2 4

	// 中断与恢复​：中间为常量显式赋值后，iota后续的值会继续递增。
	const (
		a1 = iota
		b1 = 100
		b2 = 200
		c1 = iota // 3
		d1 = iota // 4
		//  这里特别要注意，是3 4的原因。因为iota递增规律是，只要const中递增一行，iota就会+1.
		// 可以理解系统的一个全局计数器，当编译器遇到const的时候，iota默认增加1
	)
	fmt.Println(a1, b1, b2, c1, d1) // 0 100 200 3 4

	// 同一行多个常量​：同一行的多个常量使用相同的 iota值
	const (
		aa, bb = iota + 1, iota + 2 // 同一行，iota的值一样都是 = 0
		cc, dd = iota + 1, iota + 2 // 同一行, iota = 1
	)
	fmt.Println(aa, bb, cc, dd) // 1 2 2 3
}
