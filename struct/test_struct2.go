package main

import "fmt"

// 继承

// 属性首字母大写能够在包外面访问, 小写只在包内部访问，方法同理

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan继承了Human类的方法

	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()

	s := SuperMan{Human{"li4", "female"}, 88}

	s.Walk()
	s.Eat()
	s.Fly()
}
