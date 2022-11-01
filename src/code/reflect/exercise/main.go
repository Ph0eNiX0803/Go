package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int "bvhdfj"
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Println("调用成功", name)
}
func getEle(b interface{}) {
	rTy := reflect.TypeOf(b)
	rVa := reflect.ValueOf(b)
	nF := rVa.NumField()
	for i := 0; i < nF; i++ {
		fmt.Println(rVa.Field(i), rTy.Field(i))

	}
	//nM := rVa.NumMethod()
	var params []reflect.Value
	params = append(params, reflect.ValueOf("good"))
	//params = append(params, reflect.ValueOf(20))
	rVa.Method(0).Call(params)

}
func main() {
	c := Cal{}
	getEle(c)
}
