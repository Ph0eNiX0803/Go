package main

import "fmt"

func writeChan(c chan int) {
	for i := 1; i <= 2000; i++ {
		c <- i
		//fmt.Println(i)
	}
	close(c)
}
func sum(cr chan int, cl chan int64) {
	for {
		var n, ok = <-cr
		if !ok {
			break
		}
		var n1 int64 = int64(n)
		var res, i int64 = 0, 0
		for i = 1; i <= n1; i++ {
			res += i
		}
		//println(res)
		cl <- res
	}

}
func main() {
	intChan := make(chan int, 2000)
	go writeChan(intChan)
	resChan := make(chan int64, 2000)
	for i := 1; i < 8; i++ {
		go sum(intChan, resChan)
	}
	for {
		if len(resChan) == 2000 {
			close(resChan)
			break
		}
	}
	//close(resChan)
	for i := range resChan {
		fmt.Println(i)
	}
}
