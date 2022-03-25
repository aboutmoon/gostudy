package main

import "fmt"

// 声明全局变量,方法四不能声明全局变量

// 方法一
var aa int

// 方法二
var bb int = 100

// 方法三
var cc = 100

// 变量的四种声明方式
func main() {
	// 方法一：声明一个变量, 默认值是0
	var a int
	fmt.Println(a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println(b)

	// 方法三：在初始化的时候，省略数据类型，通过值自动匹配数据的类型
	var c = 100
	fmt.Println(c)
	fmt.Printf("type=%T\n", c)

	// 方法四：最常用，省略var关键词，自动推断类型
	d := 100
	fmt.Println(d)
	fmt.Printf("type=%T\n", d)

	e := "test"
	fmt.Println(e)
	fmt.Printf("type=%T\n", e)

	f := 3.14
	fmt.Println(f)
	fmt.Printf("type=%T\n", f)

	// 全局变量
	fmt.Printf("%d %d %d\n", aa, bb, cc)

	// 声明多个变量
	var xx, yy = 100, 200
	fmt.Printf("%d %d\n", xx, yy)

	// 多行多变量声明
	var (
		aa int  = 1
		bb bool = false
	)
	fmt.Printf("%d %t\n", aa, bb)

}
