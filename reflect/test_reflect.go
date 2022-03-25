package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `json:"name" info:"name" doc:"我的名字"`
	Sex  string `json:"sex" info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str)
	fmt.Printf("%+v", t)

	//for i := 0; i < t.NumField(); i++ {
	//	tagInfo := t.Field(i).Tag.Get("info")
	//	tagDoc := t.Field(i).Tag.Get("doc")
	//
	//	println(tagInfo, tagDoc)
	//}
}

func main() {
	var re resume
	findTag(&re)
}
