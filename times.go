package main

import "fmt"
import "time"

func main() {
	x := time.Now()
	y := time.Now().Weekday()
	fmt.Println(x)
	fmt.Println(y)

}
