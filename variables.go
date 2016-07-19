package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var c [5]float64
	fmt.Println("Enter a number")
	c[2] = 234
	x := make([]float64, 5, 10)
	x[0] = 32
	arr := [15]float64{12, 43, 3, 21, 5, 23, 53}
	v := arr[0:5]
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output, " ", c[2], "\n", x, "\n", arr, "\n", v)

}
