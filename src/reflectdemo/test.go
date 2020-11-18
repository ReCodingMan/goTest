package main

import (
	"fmt"
	"reflect"
)

func reflect03(b interface{}) {

	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	rTyp := rVal.Type()
	fmt.Println("rTyp=", rTyp)

	rKid := rVal.Kind()
	fmt.Println("rKid=", rKid)

	rV := rVal.Interface()
	fmt.Println("rV=", rV)

	hy := rV.(float64)
	fmt.Println("hy=", hy)
}

func main() {
	var v float64 = 1.2
	reflect03(v)
}
