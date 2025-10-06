package main

import "fmt"

func SwitchTutorial() {
	score := 80
	switch score {
	case 90:
		fmt.Println("优秀")
	case 80:
		fmt.Println("良好")
	case 70:
		fmt.Println("一般")
	default:
		fmt.Println("不及格")
	}

	// 🔹 注意：这里 switch 后没有变量，表示 case 后是 布尔表达式。
	num := 10
	switch {
	case num > 10:
		fmt.Println("大于 10")
	case num == 10:
		fmt.Println("等于 10")
	default:
		fmt.Println("小于 10")
	}

	// switch中的case支持多个值
	day := "Saturday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("工作日")
	case "Saturday", "Sunday":
		fmt.Println("周末")
	default:
		fmt.Println("未知")
	}

	/*
	1. switch 变量的基本要求
	•	变量必须是可比较的（支持 == 和 != 操作）
	•	变量的类型必须与 case 的类型一致
	•	switch 语句不会自动 fallthrough，除非显式使用
	*/
}