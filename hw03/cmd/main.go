package main

import (
	"fmt"

	chessboard "github.com/Deron-D/otus-go-basic-2025-09/hw03/pkg/chessboard"
)

func main() {
	var size int

	fmt.Print("Введите размер шахматной доски: ")

	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Ошибка: пожалуйста, введите целое число!")
		return
	}

	if size <= 0 {
		fmt.Println("Размер должен быть положительным числом!")
		return
	}

	chessboard := chessboard.CreateChessboard(size)
	fmt.Printf("\nШахматная доска %dx%d:\n", size, size)
	fmt.Println(chessboard)
}
