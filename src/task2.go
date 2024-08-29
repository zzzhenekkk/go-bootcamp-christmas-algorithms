package main

import (
	"container/heap"
	"errors"
	"fmt"
)

// Present represents a gift with a value and size.
type Present struct {
	Value int
	Size  int
}

// PresentHeap is a max-heap of presents.
type PresentHeap []Present

func (h PresentHeap) Len() int { return len(h) }

func (h PresentHeap) Less(i, j int) bool {
	if h[i].Value != h[j].Value {
		return h[i].Value > h[j].Value
	}
	return h[i].Size < h[j].Size
}

func (h PresentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PresentHeap) Push(x interface{}) {
	*h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// getNCoolestPresents returns the n coolest presents from a slice of presents.
func getNCoolestPresents(presents []Present, n int) ([]Present, error) {
	if n > len(presents) || n < 0 {
		return nil, errors.New("n is not correct")
	}

	var coolestHeap PresentHeap
	heap.Init(&coolestHeap)
	for _, present := range presents {
		heap.Push(&coolestHeap, present)
	}

	coolest := make([]Present, 0, n)
	for i := 0; i < n; i++ {
		if coolestHeap.Len() == 0 {
			break
		}
		coolest = append(coolest, heap.Pop(&coolestHeap).(Present))
	}

	return coolest, nil
}

// grabPresents returns the most valuable presents that fit into the given capacity.
func grabPresents(presents []Present, capacity int) []Present {
	n := len(presents)
	sortedPresents, err := getNCoolestPresents(presents, n)
	if err != nil {
		fmt.Println("Error sorting presents:", err)
		return nil
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if sortedPresents[i-1].Size <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-sortedPresents[i-1].Size]+sortedPresents[i-1].Value)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	res := []Present{}
	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			res = append(res, sortedPresents[i-1])
			w -= sortedPresents[i-1].Size
		}
	}

	return res
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test data
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	// Testing getNCoolestPresents function
	n := 2
	coolestPresents, err := getNCoolestPresents(presents, n)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Coolest Presents:")
		for _, present := range coolestPresents {
			fmt.Printf("Value: %d, Size: %d\n", present.Value, present.Size)
		}
	}

	// Testing grabPresents function
	capacity := 5
	selectedPresents := grabPresents(presents, capacity)
	fmt.Println("Selected Presents for Capacity", capacity, ":")
	for _, present := range selectedPresents {
		fmt.Printf("Value: %d, Size: %d\n", present.Value, present.Size)
	}
}
