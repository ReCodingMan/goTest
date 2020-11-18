package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {

	rVal := reflect.ValueOf(b)
	//看看 rVal 的kind
	fmt.Println("rVal.Kind=", rVal.Kind())
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)
}
