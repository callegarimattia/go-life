package model

// Board represents a 2D grid of cells.
// A cell is either alive (true) or dead (false).
type Board [][]bool

// Seed sets a cell to be alive.
func (b Board) Seed(x, y int) {
	if !b.checkInside(x, y) {
		return
	}

	b[y][x] = true
}

// Kill sets a cell to be dead.
func (b Board) Kill(x, y int) {
	if !b.checkInside(x, y) {
		return
	}

	b[y][x] = false
}

func (b Board) checkInside(x, y int) bool {
	return x >= 0 && x < len(b[0]) && y >= 0 && y < len(b)
}

// NewBoard creates a new board with the given width and height.
func NewBoard(width, height int) Board {
	board := make(Board, height)
	for i := range board {
		board[i] = make([]bool, width)
	}

	return board
}

// Thick returns the next state of the board by applying the rules of the game.
// Rules of the game of life by John Conway:
// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// 2. Any live cell with two or three live neighbours lives on to the next generation.
// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func (b Board) Tick() Board {
	next := NewBoard(len(b[0]), len(b))

	for y := range b {
		for x := range b[y] {
			alive := b[y][x]

			neighbours := b.countNeighbours(x, y)

			if alive {
				next[y][x] = neighbours == 2 || neighbours == 3
			} else {
				next[y][x] = neighbours == 3
			}
		}
	}

	return next
}

// countNeighbours returns the number of alive neighbours of a cell.
func (b Board) countNeighbours(x, y int) int {
	count := 0
	width := len(b[0])
	height := len(b)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx := x + dy
			ny := y + dx

			if nx >= 0 && nx < width && ny >= 0 && ny < height && b[ny][nx] {
				count++
			}
		}
	}

	return count
}
