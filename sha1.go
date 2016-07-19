package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})

	fmt.Printf("%T\n", bs)
	fmt.Println(bs)
}
