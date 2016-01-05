/* Although golang already contains priority queue and
implementing heap on top of that is trivial, I wanted to
do this for practice" */

package main

import "fmt"

type heap struct {
	arr  []int
	size int
}

func NewHeap(size int) *heap {
	return &heap{arr: make([]int, size), size: 0}
}

func (h *heap) heapify(arr []int) {
	for i := h.size - 1; i <= h.size/2; i-- {
		h.heapify_index(i)
	}
}

func (h *heap) heapify_index(i int) {
	if i <= 0 {
		return
	} else {
		if h.arr[(i+1)/2-1] > h.arr[i] {
			h.arr[i], h.arr[(i+1)/2-1] = h.arr[(i+1)/2-1], h.arr[i]
		}
	}
	h.heapify_index((i+1)/2 - 1)
}

//func (h *heap) size() int {
//	return h.size
//}

func (h *heap) top() int {
	return h.arr[0]
}

func (h *heap) pop() {
	h.arr[0] = h.arr[h.size]
	h.size--
}

func (h *heap) insert(element int) {
	fmt.Println("Inserted ", element, " at position ", h.size)
	h.arr[h.size] = element
	//append(h.arr, element)
	h.heapify_index(h.size)
	h.size++
}

func main() {
	h := NewHeap(5)
	h.insert(5)
	h.insert(6)
	h.insert(3)
	h.insert(1)
	fmt.Println(h.top())
}
