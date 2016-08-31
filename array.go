package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "soelr"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fob := [6]int{2, 2, 4, 6, 10}
	fmt.Println(fob)
}
