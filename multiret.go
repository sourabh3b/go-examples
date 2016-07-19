package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("helllo", "worlds")
	fmt.Println(a, b)
}
