package main

import (
	"encoding/json"
	"fmt"
)

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
func StructTutorial() {
	// init
	book := Books{"Go 语言", "www.runoob.com", "Go 语言教程"}
	book1 := Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程"}
	fmt.Println(book, book1)

	// init
	var book3 Books
	book3.title = "Go 语言"
	book3.author = "www.runoob.com"
	book3.subject = "Go 语言教程"
	fmt.Println(book3)
	// 作为函数参数
	printBook(book3)

	// 结构体指针
	// 结构体指针，是可以通过.语法直接访问的。底层实现，(*p).name 也就是编译器进行了省略
	book4 := &book3
	book4.title = "Go 语言指针"
	fmt.Println(book4)

	book5 := new(Books)
	book5.title = "Go 语言"
	book5.author = "www.runoob.com"
	book5.subject = "Go 语言教程"
	printBook(*book5)

	book6 := &Books{}
	book6.title = "Go 语言"
	book6.author = "www.runoob.com"
	book6.subject = "Go 语言教程"
	printBook(*book6)

	/*
		结构体方法
	*/
	p1 := Person{name: "David", age: 12, sex: "1"}
	p1.printInfo()

	p1.SetInfo("Lucy", 18)
	p1.printInfo()

	/*
		自定义类型添加方法
	*/
	var mInt myInt = 12
	mInt.printInfo()
}

type Books struct {
	title   string
	author  string
	subject string
}

func printBook(book Books) {
	fmt.Printf("%#v\n", book)
}

/*
结构体方法和接受者。
func (接受者 接受者类型) 方法名(参数列表) (返回值列表) {}
*/
type Person struct {
	name string
	age  int
	sex  string
}

func (p Person) printInfo() {
	fmt.Printf("%#v\n", p)
}

// 指针类型，一般用于方法内部修改结构体
func (p *Person) SetInfo(name string, age int) {
	p.name = name
	p.age = age
}

// 自定义类型添加方法。不能给其它包中增加方法，比如int。
type myInt int

func (m myInt) printInfo() {
	fmt.Println(m)
}

/*
结构体嵌套
*/
type Person1 struct {
	name  string
	age   int
	hobby []string

	// Person结构体中嵌套User
	user User
	// 其那套匿名结构体
	Address
}

type User struct {
	UserName string
	Password string
	Sex      string
	Aage     int
}

type Address struct {
	Location string
}

func (p *Person1) printInfo() {
	p.name = "David"
	// 注意两个写法的区别
	p.user.Aage = 12
	// 结构体嵌套
	p.Address.Location = "北京"
	p.Location = "北京2"
}

/*
结构体继承
*/
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%#v在运动", a)
}

// 子结构体
type Dog struct {
	Age    int `json:"age"` // 结构体标签
	Animal     // 结构体嵌套， 继承
}

func (d Dog) wang() {
	fmt.Printf("%#v在旺旺", d)
}

type Cat struct {
	Age     int
	*Animal // 结构体匿名嵌套， 继承指针
}

func (c Cat) miao() {
	fmt.Printf("%#v在喵喵", c)
}

func StructTutorial2() {
	d := Dog{
		Age: 12,
		Animal: Animal{
			Name: "阿奇",
		},
	}
	d.run()
	d.wang()

	c := Cat{
		Age: 12,
		Animal: &Animal{
			Name: "lili",
		},
	}
	c.run()
	c.miao()
}

/*
序列话。 "encoding/json"
*/
func StructTutorial3() {

	d := Dog{
		Age: 12,
		Animal: Animal{
			Name: "小黄",
		},
	}

	jsonByte, error := json.Marshal(d)
	fmt.Println(jsonByte, error)

	jsonstring := string(jsonByte)
	fmt.Println(jsonstring)

	str := `{"Age":12,"Name":"小黄"}`
	var d1 Dog
	err := json.Unmarshal([]byte(str), &d1)
	fmt.Println(err)
}

/*
结构体标签
*/
