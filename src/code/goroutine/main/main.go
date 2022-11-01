package main

import (
	"fmt"
	"sync"
	"time"
)

// 全局变量加锁
var (
	myMap = make(map[int64]int64, 10)
	lock  sync.Mutex
)

func test(n int64) {
	var res, i int64 = 1, 1
	for i = 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	var i int64 = 1
	for i = 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Println(i, v)
	}
	lock.Unlock()
}
