package main

import (
	"fmt"
)

//struct which holds elements
type MaxHeap struct {
	array []int
}

//add element to heap
func (h *MaxHeap) insert(key int) {
	//insert to last of array then rearrange heap
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i - 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	fmt.Println("Hello, playground")

	heap := []int{10, 1, 20, 50, 300}
	m := &MaxHeap{}
	fmt.Println(m)

	for _, val := range heap {
		m.insert(val)
		fmt.Println(m)
	}
}
