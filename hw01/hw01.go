// Программа на языке Go, которая проверяет, содержит ли массив чисел дубликаты
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 7} // Пример: содержит дубликат 1
	result := containsDuplicate(nums)
	fmt.Println(result) // Выведет: true
}
