package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Image struct {
	Id       int    `json: "id"`
	Md5      string `json: "md5"`
	Time     string `json: "time"`
	location string `json: "location"`
}

func main() {

	t := time.Now().Format(time.RFC1123)
	fmt.Println("time now : ", t)
	img := Image{
		Id:       10,
		Md5:      "kubsdjbjs32",
		Time:     t,
		location: "http://www.google.com",
	}
	// img2 := Image{
	// 	Id:       11,
	// 	Md5:      "jkbsdfk",
	// 	Time:     t,
	// 	location: "http://www.tekion.com",
	// }

	// img3 := Image{
	// 	Id:       12,
	// 	Md5:      "jksdkhb4bsdfk",
	// 	Time:     t,
	// 	location: "http://www.yahoo.com",
	// }
	// Create JSON from the instance data.

	b, _ := json.MarshalIndent(img, "", " ")
	// c, _ := json.MarshalIndent(img2, "", " ")
	// d, _ := json.MarshalIndent(img3, "", " ")
	// Convert bytes to string.
	//s := string(b)
	//fmt.Println(s)
	//fmt.Println((c))
	os.Stdout.Write(b)

	//fmt.Println(string(d))
}
