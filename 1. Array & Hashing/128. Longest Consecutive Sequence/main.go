package main

type Heap struct {
	core []int
}

func NewHeap(nums []int) *Heap {
	n := len(nums)
	h := &Heap{nums}
	if n == 0 {
		return h
	}

	for i := (n / 2) - 1; i >= 0; i-- {
		h.HeapifyDown(i)
	}
	return h
}

func (h *Heap) Pop() (int, bool) {
	n := len(h.core)
	if n == 0 {
		return 0, false
	}
	root := h.core[0]
	h.core[0] = h.core[n-1]
	h.core = h.core[:n-1]

	if n-1 > 0 {
		h.HeapifyDown(0)
	}

	return root, true
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

func longestConsecutive(nums []int) int {
	h := NewHeap(nums)
	var cont, latestCont, before int
	n, ok := h.Pop()
	if !ok {
		return 0
	}
	before = n
	cont = 1

	for {
		n, ok := h.Pop()
		if !ok {
			break
		}
		if before == n {
			continue
		}
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
