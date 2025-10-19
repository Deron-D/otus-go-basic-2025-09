package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&size)

	if size <= 0 {
		fmt.Println("Размер должен быть положительным числом!")
		return
	}

	chessboard := createChessboard(size)
	fmt.Printf("\nШахматная доска %dx%d:\n", size, size)
	fmt.Println(chessboard)
}
