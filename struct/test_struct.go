package main

import "fmt"

// 属性首字母大写能够在包外面访问, 小写只在包内部访问，方法同理

type Hero struct {
	Name  string
	Ad    int
	Level int
}

// func (this Hero) Show() {
// 	fmt.Println("Name = ", this.Name)
// 	fmt.Println("Ad = ", this.Ad)
// 	fmt.Println("Level = ", this.Level)
// }

// func (this Hero) GetName() string {
// 	return this.Name
// }

// func (this Hero) SetName(newName string) {
// 	// this 是调用该方法的对象的一个拷贝
// 	this.Name = newName
// }

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	// this 是调用该方法的对象的一个拷贝
	this.Name = newName
}

func main() {
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
