package main

import (
	"fmt"
)

/*
if expression {

} else {

}
其中expression必须是布尔表达式，即结果为true或false的表达式。
也可以是一个简单的表达式，比如a > b，a == b等。如果是表达式，需要用冒号分隔。

同时if也是有 &&，||，！等逻辑运算符的。

	switch expr {
		case case1:
			statement1
		case case2:
			statement2
		default:
			default statement
	}

switch语句中的expr表示表达式，case1，case2等表示表达式的值，statement1，statement2等表示表达式的执行语句。
同时也有fallthrough关键字，表示继续执行下一个case的语句。

也有多条件的switch语句，比如：

	switch expr {
		case case1, case2:
			statement1
		case case3, case4:
			statement2
		default:
			default statement
	}

当Switch语句中的表达式为空时，表示true，即：
	switch {

		case case1:
			statement1
		case case2:
			statement2
		default:
			default statement
	}


类型判断
	switch x.(type) {
		case type1:
			statement1
		case type2:
			statement2
		default:
			default statement
	}


Label也可以用于调整，配合goto使用。
一般很少使用goto，因为它会使代码难以理解和维护。
*/
func test() {
	a, b := 1, 2
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b")
	}

	if x := 1 + 1; x > 0 {
		fmt.Println("x > 0")
	}

	if a > b && a > 0 {
		fmt.Println("a > b && a > 0")
	}

	if a > b || a > 0 {
		fmt.Println("a > b || a > 0")
	}
}

func test2() {
	a := 1
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
		fallthrough
	default:
		fmt.Println("a == default")
	}

	switch a {
	case 1, 2:
		fmt.Println("a == 1 || a == 2")
	case 3, 4:
		fmt.Println("a == 3 || a == 4")
	default:
		fmt.Println("a == default")
	}

	switch {
	case a > 0:
		fmt.Println("a > 0")
	case a < 0:
		fmt.Println("a < 0")
	}

// label也可以用于调整，配合goto使用
   if a == 1 {
      goto A
   } else {
      fmt.Println("b")
   }
A:
   fmt.Println("a")
}
