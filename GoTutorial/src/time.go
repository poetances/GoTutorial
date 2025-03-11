package main

import (
	"fmt"
	"time"
)

func TimeTutorial() {
	timeObj := time.Now()
	fmt.Println(timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Println(year, month, day, hour, minute, second)

	// 格式化时间：2006 01 02 03 04
	fmt.Println(timeObj.Format("2006/01/02 15:04:01"))

	// 时间戳
	unixtime := timeObj.Unix()
	fmt.Println(unixtime)

	// 时间戳转时间
	timeObj2 := time.Unix(unixtime, 0)
	fmt.Println(timeObj2)
}

func TimeTickerTutorial() {
	ticker := time.NewTicker(time.Second)
	n := 5

	for t := range ticker.C {
		fmt.Println(t)
		n--
		if n == 0 {
			ticker.Stop()
			break
		}
	}
}

func TimeSleepTutorial() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	time.Sleep(time.Second)
	fmt.Println("5")
	fmt.Println("6")
}