package main

import "fmt"

// 数组传参
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
}

func main() {
	// 定义一个固定长度的数组
	var myArray1 [10]int

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	// 定义一个固定长度的数组并赋值
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	printArray(myArray3)

}
