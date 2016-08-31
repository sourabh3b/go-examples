package main

import (
	"encoding/json"
	"fmt"
)

type Informaion struct {
	Id   int
	Name []string
}

func main() {
	x := "S"
	aa, _ := json.Marshal(x)
	var s string = string(aa)
	fmt.Printf("%T\n", s)
	fmt.Println(s)

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	sliceA := []string{"aaa", "bbb", "ccc"}
	slcA, _ := json.Marshal(sliceA)
	fmt.Println(string(slcA))

	mapA := map[string]int{"apple": 12, "banana": 33, "peach": 43}
	mA, _ := json.Marshal(mapA)
	fmt.Println(string(mA))

}
