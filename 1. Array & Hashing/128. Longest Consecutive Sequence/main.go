package main

import (
	"fmt"
	"io"
)

type Heap struct {
	core []int
}

func (h *Heap) Push(t int) {
	h.core = append(h.core, t)
	h.HeapifyUp(len(h.core) - 1)
}

func (h *Heap) Pop() (int, error) {
	n := len(h.core)
	if n == 0 {
		return 0, io.EOF
	}
	root := h.core[0]
	h.core[0] = h.core[n-1]
	h.core = h.core[:n-1]

	if n-1 > 0 {
		h.HeapifyDown(0)
	}

	return root, nil
}

func (h *Heap) HeapifyDown(index int) {
	n := len(h.core)
	left := 2*index + 1
	right := 2*index + 2
	smallest := index

	if left < n && h.core[left] < h.core[smallest] {
		smallest = left
	}

	if right < n && h.core[right] < h.core[smallest] {
		smallest = right
	}
	if smallest != index {
		h.core[smallest], h.core[index] = h.core[index], h.core[smallest]
		h.HeapifyDown(smallest)
	}
}

func (h *Heap) HeapifyUp(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2

	if h.core[index] < h.core[parent] {
		h.core[index], h.core[parent] = h.core[parent], h.core[index]
		h.HeapifyUp(parent)
	}
}

func longestConsecutive(nums []int) int {
	var h Heap
	for i := range nums {
		h.Push(nums[i])
	}

	var cont, latestCont, before int
	n, err := h.Pop()
	if err == io.EOF {
		return 0
	}
	before = n
	cont = 1

	for {
		n, err := h.Pop()
		if err == io.EOF {
			break
		}
		if before == n {
			continue
		}
		fmt.Println(before, n, cont, latestCont)
		if before-n == -1 || before-n == 1 {
			cont += 1
		} else {
			if cont > latestCont {
				latestCont = cont
			}
			cont = 1
		}
		before = n
	}

	if cont > latestCont {
		return cont
	}
	return latestCont
}
