package main

import (
	"fmt"
	"math"
)

func write(c chan int) {
	for i := 1; i < 80000000; i++ {
		c <- i
	}
	close(c)

}
func isPrime(i, o, e chan int) {

	for {
		//time.Sleep(time.Millisecond)
		n, ok := <-i
		if !ok {
			break
		}
		flag := true
		//fmt.Println(n)
		for k := 2; k <= int(math.Sqrt(float64(n))); k++ {
			if n%k == 0 {
				flag = false
				break
			}
		}
		if flag {
			o <- n
		}
	}
	fmt.Println("取不到数据，协程退出")
	e <- 1
}
func main() {
	intChan := make(chan int, 1000)
	resChan := make(chan int, 80000000)
	exitChan := make(chan int, 12)
	go write(intChan)
	for i := 1; i <= 12; i++ {
		go isPrime(intChan, resChan, exitChan)
	}
	for {
		//fmt.Println(len(exitChan))
		if len(exitChan) == 12 {

			close(exitChan)
			close(resChan)
			break
		}
	}
	for i := range resChan {
		fmt.Println(i)
	}
}
