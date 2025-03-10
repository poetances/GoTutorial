package main

import "fmt"

/*
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}

variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
*/
func StructTutorial() (string, int) {
	book := Books{"Go 语言", "www.runoob.com", "Go 语言教程"}
	book1 := Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程"}
	fmt.Println(book, book1)

	var book3 Books
	book3.title = "Go 语言"
	book3.author = "www.runoob.com"
	book3.subject = "Go 语言教程"
	fmt.Println(book3)

	// 作为函数参数
	printBook(book3)

	// 结构体指针
	book4 := &book3
	book4.title = "Go 语言指针"
	fmt.Println(book4)
	return "hello", 1
}

type Books struct {
	title string
	author string
	subject string
}

func printBook(book Books) {
	fmt.Println(book)
}