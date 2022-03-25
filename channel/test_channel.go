package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end")

		println("goroutine running ...")

		// 如果channel中的值不被读取，就会被阻塞
		c <- 666
	}()

	time.Sleep(5 * time.Second)

	// 如果从channel中没有值，会被阻塞
	num := <-c

	fmt.Printf("num = %d \n", num)
	fmt.Println("main goroutine end")
}
