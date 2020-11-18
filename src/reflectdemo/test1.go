package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{})  {
	//通过反射，获取变量的 type kind 值

	//1、先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	//2、获取 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	fmt.Printf("rVal=%v, rVal-type=%T\n", rVal, rVal)

	//3、将 rVal 转成 interface{}
	iv := rVal.Interface()
	//将 Interface{} 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=", num2)

}

func main()  {

	var num int = 100
	reflectTest01(num)
}


