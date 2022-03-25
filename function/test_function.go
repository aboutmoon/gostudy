package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)

	c := 10

	return c
}

// 匿名返回值
func foo2(a string, b int) (string, int) {
	fmt.Println(a)
	fmt.Println(b)

	c := 10
	d := "string"

	return d, c
}

// 返回值带名称
func foo3(a string, b int) (c string, d int) {
	fmt.Println(a)
	fmt.Println(b)

	d = 10
	c = "string"

	return c, d
}

func main() {
	c := foo1("test", 100)
	fmt.Println(c)

	d, e := foo2("test2", 200)
	fmt.Println(d, e)
}
