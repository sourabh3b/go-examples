package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	r, err := regexp.Compile(`hello`)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	if r.MatchString("hllo hwllo") == true {
		fmt.Println("Match")
	} else {
		fmt.Println("Mis match")
	}

}
