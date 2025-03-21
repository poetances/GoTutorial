package main

import "fmt"

func main() {
	ch := make(chan string) // 创建 channel
	go func() {
		ch <- "Hello Goroutine" // 发送数据
	}()

	msg := <-ch // 接收数据
	fmt.Println(msg)
}
