package main

import "fmt"

type MaxHeap struct {
	array []int
}

// Insert key to the last index and rearange MaxHeap
func (m *MaxHeap) Insert(key int) {
	m.array = append(m.array, key)
	m.maxHeapifyUp(len(m.array) - 1)
}

// Returning the lasgest key , and romove it from the heap
func (m *MaxHeap) Extract() int {
	extracted := m.array[0]

	// When the array is empty
	if len(m.array) == 0 {
		fmt.Println("Cannot extract heap because array is empty")
		return -1
	}

	lastIndex := len(m.array) - 1
	m.array[0] = m.array[lastIndex]
	m.array = m.array[:lastIndex]

	m.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.array[parent(index)] < m.array[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyUp will heapify from top to bottom
func (m *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(m.array) - 1
	l, r := left(index), right(index)
	childIndexToCompare := 0

	for l <= lastIndex {
		if l == lastIndex { // When left child is the only child
			childIndexToCompare = l
		} else if m.array[l] > m.array[r] { // when left child is larger than right child
			childIndexToCompare = l
		} else { // when right child is larger than left child
			childIndexToCompare = r
		}

		// Compare array value of current index to larger child and swap it smaller
		if m.array[index] < m.array[childIndexToCompare] {
			m.swap(index, childIndexToCompare)
			index = childIndexToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

	for m.array[parent(index)] < m.array[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

// Get the parent index
func parent(index int) int {
	return (index - 1) / 2
}

// Get the left child index
func left(index int) int {
	return index*2 + 1
}

// Get the left child index
func right(index int) int {
	return index*2 + 2
}

// Swap 2 value of heap array with their index
func (m *MaxHeap) swap(i1 int, i2 int) {
	m.array[i1], m.array[i2] = m.array[i2], m.array[i1]
}

func main() {
	m := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 23, 21, 31, 32, 324, 4, 234}
	for _, item := range buildHeap {
		m.Insert(item)
		fmt.Println(m)
	}

	fmt.Println("\n===============================")

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
