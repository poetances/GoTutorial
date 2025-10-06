package main

import "fmt"


func BasicDataTutorial() {

	var i int
	var f float64
	var b bool
	var c complex64
	var bb byte	= 0xff

	fmt.Println(i, f, b, c, bb)

	var c1 complex64 = 1 + 2i
	fmt.Println(c1)
}