package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c, d = iota, iota
	)

	fmt.Println(a,b,c,d)
}
