package main

import (
	"fmt"
)

/*
Golagng中接口定义：

	type 接口名 interface {
		方法名1（参数列表1）返回值列表1
		方法名2（参数列表2）返回值列表2
		....
	}

接口也是一种类型
*/
func InterfaceTutorial() {
	p := Phone{
		Name: "华为手机",
	}

	var p1 Usber = p
	p1.start()
	p1.stop()

	c := Camera{"索尼相机"}
	var c1 Usber = c
	c1.start()
	c1.stop()

	c.run()
	c.ptrRun()
}

type A interface{} // 空接口 表示没有任何约束、任意类型都可以实现空接口
// 空接口
func InterfaceTutorial1() {
	var a A
	var str = "你好golang"
	a = str // 让字符串实现A这个接口
	fmt.Printf("值：%v 类型：%T", a, a)

	// golang中空接口也可以是一种类型。可以表示任意类型
	var a1 interface{} = 12
	fmt.Println(a1)

	// 常见用法> interface{}类似Any，表示字典的value可以是任意类型
	var m1 = make(map[string]interface{})
	m1["name"] = 12
	m1["age"] = "david"

	var s1 = []interface{}{1, "2", true}
	fmt.Println(s1...)

	/*
		下面是Println方法
		// any is an alias for interface{} and is equivalent to interface{} in all ways.
		// type any = interface{}
		func Println(a ...any) (n int, err error) {
			return Fprintln(os.Stdout, a...)
		}
	*/

	// 类型断言。因为空接口可以是任意类型，那么这个时候就需要类型断言，其语法格式：x.(T)
	v, ok := a1.(string)
	fmt.Println(v, ok)

	v1, ok1 := a1.(int) // 如果断言成功，会将值赋值给v1
	fmt.Println(v1, ok1)
}

func InterfacePrint(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println(v, "string类型")
	}
}

func InterfacePrint1(x interface{}) {
	// 这个语法有限制，只能用于swiftch语法
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	}
}

type Usber interface {
	start()
	stop()
}

type Battery interface {
	recharge()
}

// 接口嵌套
type Computer interface {
	Usber
	Battery
}

// 如果接口里有方法的话，必须通过结构体或者自定义类型实现接口
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
	Name string
}

func (c Camera) start() {
	fmt.Println(c.Name, "启动")
}

func (c Camera) stop() {
	fmt.Println(c.Name, "关机")
}

// 结构体方法
func (c Camera) run() {
	fmt.Println(c.Name, "运行")
}
func (c *Camera) ptrRun() {
	fmt.Println(c.Name, "运行")
}