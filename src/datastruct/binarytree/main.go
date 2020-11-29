package main

type Hero struct {
	No int
	Name string
	Left *Hero
	Right *Hero
}

func main() {
	//构建一个二叉树
	root := &Hero{
		No: 1,
		Name: "宋江",
	}

	left1 := &Hero{
		No: 2,
		Name: "吴用",
	}
}
