package main

import "fmt"

func main() {
	// 第一种声明方式,声明一个变量，只是一个指针
	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("myMap1 is empty")
	}

	// 在使用map前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["4"] = "python"
	myMap1["5"] = "python"
	myMap1["6"] = "python"
	myMap1["7"] = "python"
	myMap1["8"] = "python"
	myMap1["9"] = "python"
	myMap1["10"] = "python"
	myMap1["11"] = "python"
	myMap1["12"] = "python"
	myMap1["13"] = "python"
	myMap1["14"] = "python"
	myMap1["15"] = "python"
	myMap1["16"] = "python"

	fmt.Println(myMap1)

	// 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	// 第三种声明方式
	myMap3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}

	fmt.Println(myMap3)
}
