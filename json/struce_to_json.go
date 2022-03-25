package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal err", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	myMovie := Movie{}

	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal err", err)
	}

	fmt.Printf("%v \n", myMovie)

}
