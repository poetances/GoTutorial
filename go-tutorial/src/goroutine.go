package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 用于同步主线程和携程的执行
var wg sync.WaitGroup
/*
 goroutine，go中的协程。
*/
func GoroutineTutorial() {
	cpunum := runtime.NumCPU()
	fmt.Println("CPU个数", cpunum)

	// 设置CPU使用的数量
	runtime.GOMAXPROCS(cpunum - 1)

	wg.Add(1) // 协程 + 1
	// 开启一个携程
	go goroutine()

	for i := 0; i < 10; i++ {
		fmt.Println("main golang")
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait() // 等待协程执行完毕
}

func goroutine() {
	for i := 0; i < 10; i++ {
		fmt.Println("goroutine golang")
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程-1
}