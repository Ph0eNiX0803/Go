package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
	Age  int
}

func test01(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)
	rVal := reflect.ValueOf(b)
	iv := rVal.Interface()
	i := iv.(int)
	fmt.Println(i)
	fmt.Println(rVal)
}

func test02(b interface{}) {

}
func getFloat64(b interface{}) {
	rVal := reflect.ValueOf(b)
	rTyp := reflect.TypeOf(b)
	rKind := reflect.ValueOf(b).Kind()
	fmt.Println(rVal, rTyp, rKind)
	rInterface := rVal.Interface()
	rV := rInterface.(float64)
	fmt.Println(rV)
}
func main() {
	var v float64 = 1.2
	getFloat64(v)
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("%v\n", str)
	//num := 100
	//test01(num)
}
