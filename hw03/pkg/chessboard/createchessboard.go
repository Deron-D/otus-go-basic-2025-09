package chessboard

import (
	"strings"
)

func createChessboard(size int) string {
	var board strings.Builder

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board.WriteString("#")
			} else {
				board.WriteString(" ")
			}
		}
		if i < size-1 {
			board.WriteString("\n")
		}
	}

	return board.String()
}
