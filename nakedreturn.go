package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum + 1
	y = sum * 7
	return
}

var a, b, c int = 1, 2, 3

//  k:=3 //This declaration cannot be called outside the main

func main() {
	var i int
	var e, f, g = false, true, "no!"
	fmt.Println(split(2))
	fmt.Println(i)
	fmt.Println(a, b, c, e, f, g)
	k := 2 //since k :=2 is defined inside main , so is can be used only in this func
	fmt.Println(k)
}
