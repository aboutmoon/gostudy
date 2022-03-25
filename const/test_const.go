package main

import "fmt"

// const来定义枚举类型
// iota 每行的iota+1，第一行默认值0
const (
	BEIJING = 10 * iota
	SHANGHAI
	XIAN
)

const (
	CHINA    = 0
	AMERACAN = 1
)

const (
	a, b = iota + 1, iota + 2   // iota = 0, a = 1, b = 2
	c, d                        // iota = 1, c = 2, d= 3
	e, f                        // iota = 2, c = 3, d= 4
	g, h = iota + 10, iota + 20 // iota = 3, g = iota + 10 = 13, h = iota + 20 = 24
	i, j                        // iota = 4, i = iota + 10 = 14, j = iota + 10 = 24
)

func main() {
	// 常量（只读属性）
	const length int = 100

	// length = 20 常量不能修改

	// var a = iota iota只能配合const使用

	fmt.Println(length)

	fmt.Println(BEIJING, SHANGHAI, XIAN)

	fmt.Println(CHINA, AMERACAN)

	fmt.Println(a, b, c, d, e, f, g, h, i, j)

	fmt.Printf("%T, %v \n", a, a)
	fmt.Printf("%T, %v \n", b, b)
}
