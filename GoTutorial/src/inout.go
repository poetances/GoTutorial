package main

import (
	"fmt"
	"os"
)

func InoutTutorial()	{
	fmt.Printf("%%%s\n", "hello world")
 
	fmt.Printf("%s\n", "hello world") 
	fmt.Printf("%q\n", "hello world") 
	fmt.Printf("%d\n", 2<<7-1)			
	
	fmt.Printf("%f\n", 1e2)			
	fmt.Printf("%e\n", 1e2)				
	fmt.Printf("%E\n", 1e2)				
	fmt.Printf("%g\n", 1e2)				
	
	fmt.Printf("%b\n", 2<<7-1)			
	fmt.Printf("%#b\n", 2<<7-1)			
	fmt.Printf("%o\n", 2<<7-1)			
	fmt.Printf("%#o\n", 2<<7-1)			
	fmt.Printf("%x\n", 2<<7-1)			
	fmt.Printf("%#x\n", 2<<7-1)			
	fmt.Printf("%X\n", 2<<7-1)			
	fmt.Printf("%#X\n", 2<<7-1)			
	
	type person struct {
		name    string
		age     int
		address string
	}
	fmt.Printf("%v\n", person{"lihua", 22, "beijing"})	
	fmt.Printf("%+v\n", person{"lihua", 22, "beijing"})	
	fmt.Printf("%#v\n", person{"lihua", 22, "beijing"})	
	fmt.Printf("%t\n", true)							
	fmt.Printf("%T\n", person{})						
	fmt.Printf("%c%c\n", 20050, 20051)					
	fmt.Printf("%U\n", '码')							   
	fmt.Printf("%p\n", &person{})						

	// Sprintf有返回值
	result := fmt.Sprintf("%d", 123)
	fmt.Println(result)

	// Fprintf
	// Fprintf 用于格式化并输出到 io.Writer 接口类型的 w 中，返回写入的字节数和遇到的任何写入错误。
	// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	fmt.Fprint(os.Stdout, "hello world")

	
	// println
	// 特点
	// •	内置函数：println 是 Go 语言的一个低级别内置函数，用于直接打印到标准输出。
	// •	简单输出：仅用于快速调试，不能进行格式化操作。
	// •	无严格保证：其行为和输出格式在不同平台或版本的 Go 编译器中可能有所不同，不建议在生产代码中使用。
}