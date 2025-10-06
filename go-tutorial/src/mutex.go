package main

import (
	"fmt"
	"sync"
)

var mutext sync.Mutex
var count int = 0

func MutexTutorial() {
	mutext.Lock()
	count++
	fmt.Println(count)
	mutext.Unlock()
}
