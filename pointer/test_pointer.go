package main

import "fmt"

func changeValue(p *int) {
	*p = 10
}

func swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	i := 1

	changeValue(&i)

	fmt.Println(i)

	a := 5
	b := 10

	swap(&a, &b)

	fmt.Println(a, b)
}
