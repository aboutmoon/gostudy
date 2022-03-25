package lib2

import "fmt"

const a = "lib2a"

var b = "lib2b"

func init() {
	fmt.Println("lib2 init")
	fmt.Println(a, b)
}

func Test() {
	fmt.Println("lib2 test")
}
