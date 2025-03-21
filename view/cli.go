package view

import (
	"github.com/callegarimattia/go-life/model"
)

func Render(b model.Board) {
	print("\033[H\033[2J") // Clears the screen.

	for y := range b {
		for x := range b[y] {
			if b[y][x] {
				print("█")
			} else {
				print(" ")
			}
		}

		println()
	}

	println()
}
