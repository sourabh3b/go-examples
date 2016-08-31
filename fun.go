package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addd(x, y int) int {
	return x + y
}
func main() {
	fmt.Println("Program for add")
	fmt.Println(add(4, 5))
	fmt.Println(addd(46, 5))

}
