package main

import "fmt"

// Определение структуры Present
type Present struct {
	Value int
	Size  int
}

// Функция для решения задачи о рюкзаке
func grabPresents(presents []Present, capacity int) []Present {
	n := len(presents)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Заполнение таблицы dp
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if presents[i-1].Size <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-presents[i-1].Size]+presents[i-1].Value)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// Восстановление набора подарков
	res := []Present{}
	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			res = append(res, presents[i-1])
			w -= presents[i-1].Size
		}
	}

	return res
}

// Вспомогательная функция для нахождения максимального значения
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Пример использования
func main() {
	presents := []Present{
		{Value: 5, Size: 3},
		{Value: 6, Size: 2},
		{Value: 3, Size: 1},
	}
	capacity := 5

	result := grabPresents(presents, capacity)
	fmt.Println("Selected Presents:")
	for _, present := range result {
		fmt.Printf("Value: %d, Size: %d\n", present.Value, present.Size)
	}
}
