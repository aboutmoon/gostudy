package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	// copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3, 10) // s2 = [0,0,0]

	// 将s中的值依次拷贝到s2
	copy(s2, s)

	s1[0] = 50
	s2[2] = 10

	fmt.Println(s2, len(s2), cap(s2))
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	fmt.Println(s2, len(s2), cap(s2))
}
