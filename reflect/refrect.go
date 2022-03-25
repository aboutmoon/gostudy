package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
}

func (u *User) SayHello() {
	println("Hello")
}

type Say interface {
	SayHello()
}

func main() {
	user := User{1, "hb"}

	// 反射获取类型
	inputType := reflect.TypeOf(user)
	fmt.Println("type is :", inputType)

	inputType2 := reflect.TypeOf(&user)
	fmt.Println(inputType2.NumMethod())
	fmt.Println(inputType2.Method(0))

	for _, v := range inputType2.NumMethod() {
		p
	}

}
