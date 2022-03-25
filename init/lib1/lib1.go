package lib1

import "fmt"

const a = "lib1a"

var b = "lib1b"

func init() {
	fmt.Println("lib1 init")
	fmt.Println(a, b)
}

func Test() {
	fmt.Println("lib1 test")
}
