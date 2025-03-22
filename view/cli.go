package view

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/callegarimattia/go-life/model"
)

// Render prints the board to the given writer.
func Render(w io.Writer, b model.Board) {
	clearScreen()

	for y := range b {
		for x := range b[y] {
			if b[y][x] {
				fmt.Fprint(w, "O")
			} else {
				fmt.Fprint(w, ".")
			}
		}

		fmt.Fprintln(w)
	}

	fmt.Fprintln(w)
}

// only works on Unix-like systems.
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	_ = cmd.Run()
}
