package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		println("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//go newTask()
	//
	//i := 0
	//for {
	//	i++
	//	println("Main Goroutine : i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}

	go func() {
		defer println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
