package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Channel是一种类型，一种引用类型。声明管道的格式如下：
var 变量 chan 类型
如：
var ch1 chan int   整型管道
var ch2 chan bool  布尔类型管道
var ch3 chan []int 切片类型的管道

// make创建管道
*/
func ChannelTutorial() {

	// 创建
	ch := make(chan int, 3) // 3，表示管道的容量
	// 存储数据到管道
	ch <- 10
	ch <- 30
	ch <- 50
	// 获取管道里面的数据
	<-ch // 取一次值
	a := <-ch
	fmt.Println(a) // 30. 管道里面的数据遵循FIFO


	fmt.Printf("值：%v 容量:%v 长度：%v\n", ch, cap(ch), len(ch))


	// 管道是引用类型
	ch1 := make(chan int, 4)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	ch2 := ch1
	ch2 <- 4


	// 管道阻塞
	ch3 := make(chan int ,1)
	ch3 <- 1
	// ch3 <-2 // error: all goroutines are asleep - deadlock! 如果管道容量是1，往里面放多余1的数，将会阻塞

	a1 := <-ch3
	// a2 := <-ch3 // error: all goroutines are asleep - deadlock! 同样，如果管道没有数据，还去取，还是会报错
	fmt.Println(a1)
}

// 循环遍历管道
func ChannelTutorial1() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)// 关闭管道。不然遍历会报错fatal error: all goroutines are asleep - deadlock!

	// 遍历。在没有携程的情况下，如果管道中没有数据，还继续读取就会报错
	for v := range ch {
		fmt.Println(v)
	}

	// 通过for 循环不需要close
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

/// goroutine + chan
var ws sync.WaitGroup
func read(ch chan int) {
	for v := range ch {
		fmt.Println("读取数据：", v)
		time.Sleep(time.Second)
	}
	ws.Done()
}

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("写入数据：", i)
		time.Sleep(time.Second)
	}
	close(ch)
	ws.Done()
}

func ChannelTutorial2() {
	ch := make(chan int, 10)
	ws.Add(2)
	go read(ch)
	go write(ch)
	ws.Wait()
}