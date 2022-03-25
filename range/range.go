package main

import (
	"fmt"
	"time"
)

func range1(slice []string) {
	for index, value := range slice {
		println(index, value)
	}
}

func rangeMap(m map[int]string) {
	for key, value := range m {
		_, _ = key, value
	}
}

func rangeAppend() {
	v := []int{1, 2, 3}
	for i := range v {
		println(i)
		v = append(v, i)
	}
}

func timeMeasurement(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}

func main() {
	var a []string
	b := make(map[int]string)

	for i := 1000; i < 10000; i++ {
		a = append(a, fmt.Sprintf("%d", i))
		b[i] = fmt.Sprintf("%d", i)
	}
	//
	defer timeMeasurement(time.Now())
	//rangeMap(b)
	rangeAppend()

}
