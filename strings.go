package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "1,2,3,4,5"
	result := strings.Split(data, ",")
	fmt.Println(len(data))

	for i := range result {
		fmt.Println(result[i])
	}

}
