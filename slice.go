package main

import "fmt"

func main() {
	prime := [6]int{2, 4, 234, 453, 3}

	var s []int = prime[1:4]
	fmt.Println(s)
}
