package main

import (
	"fmt"
	// _ "go-study/init/lib1"
	. "go-study/init/lib2"
)

// _ 包名 匿名导入：无法使用当前包的方法，但是会执行该包内部的init()方法
// myLIb2 包名 给包取个别名，通过myLib2.method()使用
// . 导入所有方法，直接使用包中的方法名调用，尽量不用，防止两个包有同名方法

var i = "maini"

const j = "mainj"

func init() {
	fmt.Println("main init")
	fmt.Println(i, j)
}

func main() {
	// lib1.Test()
	Test()
}
