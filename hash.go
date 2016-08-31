package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}
func main() {
	h1, err := getHash("tes1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	// h.Write([]byte("sourabh bhagat"))
	// v := h.Sum32()
	// fmt.Println(v)
	fmt.Println(h1, h1, h1 == h2)
}
