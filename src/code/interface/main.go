package main

import (
	"fmt"
	"runtime"
)

type Usb interface {
	sum(int, int) int
	start()
	stop()
}

type Usb2 interface {
	sum()
	start()
	stop()
}

type Phone struct {
}

func (p Phone) start() {
	fmt.Println("开始工作")
}
func (p Phone) sum(i, j int) int {
	return 10
}

func (p Phone) stop() {
	fmt.Println("停止工作")
}

type Computer struct {
}

func working(usb Usb) {
	usb.start()
	usb.stop()
	fmt.Println(usb.sum(6, 6))
}
func sum() {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}
func main() {
	cpu := runtime.NumCPU()
	fmt.Println("cpuNum=", cpu)
	go sum()
	//c := Computer{}
	p := Phone{}
	working(p)
	var i interface{}
	i = p
	var p1 = Phone{}
	p1 = i.(Phone)
	fmt.Println(p1)
}
