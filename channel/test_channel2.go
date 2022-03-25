package main

import (
	"fmt"
	"time"
)

// 当channel已满，再向里面写数据会被阻塞
// 当channel为空，再从里面取数据会被阻塞
func main() {
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), ",cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("child goroutine end")

		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("child goroutine running, send ", i)
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("main end")
}
