package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{53.453, -34.03}
	fmt.Println(m["Bell Labs"])

	ma := make(map[string]int)

	ma["Answer"] = 43
	fmt.Println("Value : ", ma["Answer"])
	ma["Answer"] = 46

	delete(m, "Answer")
	fmt.Println("New val", ma["Answer"])

}
