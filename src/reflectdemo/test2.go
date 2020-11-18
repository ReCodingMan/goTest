package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest02(b interface{})  {
	//通过反射，获取变量的 type kind 值

	//1、先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)
	fmt.Println("rTyp.Kind=", rTyp.Kind())

	//2、获取 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	fmt.Println("rVal.Kind=", rVal.Kind())
	fmt.Printf("rVal=%v, rVal-type=%T\n", rVal, rVal)

	//3、将 rVal 转成 interface{}
	iv := rVal.Interface()
	fmt.Printf("iv=%v, iv-type=%T\n", iv, iv)

	//将 Interface{} 通过断言转成需要的类型
	num2 := iv.(Student)
	fmt.Println("num2=", num2)

}

type Student struct {
	Name string
	Age int
}

func main() {

	stu := Student{
		Name : "tom",
		Age : 10,
	}

	reflectTest02(stu)
}


