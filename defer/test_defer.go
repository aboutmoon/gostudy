package main

import "fmt"

// defer和return谁先谁后

func deferFunc() int {
	fmt.Println("defer func")
	return 0
}

func returnFunc() int {
	fmt.Println("return func")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()

	return returnFunc()
}

func main() {

	defer fmt.Println("main end 1")
	defer fmt.Println("main end 2")
	fmt.Println("go 1")
	fmt.Println("go 2")

	deferAndReturn()
}
