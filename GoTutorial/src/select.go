package main

import (
	"fmt"
	"time"
)

/*
1.1.1. select 语句
select 语句是 Go 语言中的多路通道（channel）选择机制，主要用于处理多个 channel 之间的并发通信。它的作用类似于 switch，但 select 只适用于 channel 操作。
select {
	case <-ch1:
		// 从 ch1 读取数据
	case ch2 <- value:
		// 向 ch2 发送数据
	default:
		// 如果没有可用的 channel，则执行 default 代码块
	}


*/
func SelectTutorial() {

	ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "来自 ch1 的消息"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "来自 ch2 的消息"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println("接收到：", msg1)
    case msg2 := <-ch2:
        fmt.Println("接收到：", msg2)
    }
}