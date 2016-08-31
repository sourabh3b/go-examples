package main

import "fmt"
import "time"

func main() {
	fmt.Println("Timeinf")

	today := time.Now().Weekday()

	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away.")

	}

	defer fmt.Println("Hewllo")
	defer fmt.Println("Hewllosa")
	fmt.Println("owl")
}
