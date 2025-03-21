package shapes

import "github.com/callegarimattia/go-life/model"

// Oscillators:
// - Blinker
// - Toad
// - Beacon

// SeedBlinker seeds a blinker oscillator.
func SeedBlinker(board model.Board, x, y int) {
	board.Seed(x, y-1)
	board.Seed(x, y)
	board.Seed(x, y+1)
}

// SeedToad seeds a toad oscillator.
func SeedToad(board model.Board, x, y int) {
	board.Seed(x+1, y)
	board.Seed(x+2, y)
	board.Seed(x+3, y)
	board.Seed(x, y+1)
	board.Seed(x+1, y+1)
	board.Seed(x+2, y+1)
}

// SeedBeacon seeds a beacon oscillator.
func SeedBeacon(board model.Board, x, y int) {
	board.Seed(x, y)
	board.Seed(x+1, y)
	board.Seed(x, y+1)
	board.Seed(x+1, y+1)

	board.Seed(x+2, y+2)
	board.Seed(x+3, y+2)
	board.Seed(x+2, y+3)
	board.Seed(x+3, y+3)
}

// Spaceships:
// - Glider
// - Lightweight Spaceship (LWSS)
// - Middleweight Spaceship (MWSS)
// - Heavyweight Spaceship (HWSS)

// SeedGlider seeds a glider spaceship.
func SeedGlider(board model.Board, x, y int) {
	board.Seed(x+1, y)
	board.Seed(x+2, y+1)
	board.Seed(x, y+2)
	board.Seed(x+1, y+2)
	board.Seed(x+2, y+2)
}

// SeedLWSS seeds a Lightweight Spaceship (LWSS).
func SeedLWSS(board model.Board, x, y int) {
	board.Seed(x, y)
	board.Seed(x+3, y)
	board.Seed(x+4, y+1)
	board.Seed(x, y+2)
	board.Seed(x+4, y+2)
	board.Seed(x+1, y+3)
	board.Seed(x+2, y+3)
	board.Seed(x+3, y+3)
	board.Seed(x+4, y+3)
}

// SeedMWSS seeds a Middleweight Spaceship (MWSS).
func SeedMWSS(board model.Board, x, y int) {
	board.Seed(x, y)
	board.Seed(x+1, y)
	board.Seed(x+2, y)
	board.Seed(x+3, y)
	board.Seed(x+4, y)
	board.Seed(x+5, y+1)
	board.Seed(x, y+2)
	board.Seed(x+5, y+2)
	board.Seed(x+1, y+3)
	board.Seed(x+2, y+3)
	board.Seed(x+3, y+3)
	board.Seed(x+4, y+3)
}

// SeedHWSS seeds a Heavyweight Spaceship (HWSS).
func SeedHWSS(board model.Board, x, y int) {
	board.Seed(x, y)
	board.Seed(x+1, y)
	board.Seed(x+2, y)
	board.Seed(x+3, y)
	board.Seed(x+4, y)
	board.Seed(x+5, y)
	board.Seed(x+6, y+1)
	board.Seed(x, y+2)
	board.Seed(x+6, y+2)
	board.Seed(x+1, y+3)
	board.Seed(x+2, y+3)
	board.Seed(x+3, y+3)
	board.Seed(x+4, y+3)
	board.Seed(x+5, y+3)
}

// Methuselahs:
// - R-pentomino
// - Diehard
// - Acorn

// SeedRpentomino seeds the R-pentomino methuselah.
func SeedRpentomino(board model.Board, x, y int) {
	board.Seed(x+1, y)
	board.Seed(x+2, y)
	board.Seed(x, y+1)
	board.Seed(x+1, y+1)
	board.Seed(x+1, y+2)
}

// SeedDiehard seeds the diehard methuselah.
func SeedDiehard(board model.Board, x, y int) {
	board.Seed(x+6, y)
	board.Seed(x, y+1)
	board.Seed(x+1, y+1)
	board.Seed(x+1, y+2)
	board.Seed(x+5, y+2)
	board.Seed(x+6, y+2)
	board.Seed(x+7, y+2)
}

// SeedAcorn seeds the acorn methuselah.
func SeedAcorn(board model.Board, x, y int) {
	board.Seed(x+1, y)
	board.Seed(x+3, y+1)
	board.Seed(x, y+2)
	board.Seed(x+1, y+2)
	board.Seed(x+4, y+2)
	board.Seed(x+5, y+2)
	board.Seed(x+6, y+2)
}

// Guns:
// - Gosper Glider Gun
// - Simkin Glider Gun

// SeedGosperGliderGun seeds the Gosper Glider Gun.
func SeedGosperGliderGun(board model.Board, x, y int) {
	board.Seed(x+24, y)
	board.Seed(x+22, y+1)
	board.Seed(x+24, y+1)
	board.Seed(x+12, y+2)
	board.Seed(x+13, y+2)
	board.Seed(x+20, y+2)
	board.Seed(x+21, y+2)
	board.Seed(x+34, y+2)
	board.Seed(x+35, y+2)
	board.Seed(x+11, y+3)
	board.Seed(x+15, y+3)
	board.Seed(x+20, y+3)
	board.Seed(x+21, y+3)
	board.Seed(x+34, y+3)
	board.Seed(x+35, y+3)
	board.Seed(x, y+4)
	board.Seed(x+1, y+4)
	board.Seed(x+10, y+4)
	board.Seed(x+16, y+4)
	board.Seed(x+20, y+4)
	board.Seed(x+21, y+4)
	board.Seed(x, y+5)
	board.Seed(x+1, y+5)
	board.Seed(x+10, y+5)
	board.Seed(x+14, y+5)
	board.Seed(x+16, y+5)
	board.Seed(x+17, y+5)
	board.Seed(x+22, y+5)
	board.Seed(x+24, y+5)
	board.Seed(x+10, y+6)
	board.Seed(x+16, y+6)
	board.Seed(x+24, y+6)
	board.Seed(x+11, y+7)
	board.Seed(x+15, y+7)
	board.Seed(x+12, y+8)
	board.Seed(x+13, y+8)
}

// SeedSimkinGliderGun seeds the Simkin Glider Gun.
func SeedSimkinGliderGun(board model.Board, x, y int) {
	board.Seed(x+1, y)
	board.Seed(x+2, y)
	board.Seed(x+1, y+1)
	board.Seed(x+2, y+1)
	board.Seed(x+13, y)
	board.Seed(x+14, y)
	board.Seed(x+13, y+1)
	board.Seed(x+14, y+1)
	board.Seed(x+12, y+2)
	board.Seed(x+15, y+2)
	board.Seed(x+11, y+3)
	board.Seed(x+16, y+3)
	board.Seed(x+11, y+4)
	board.Seed(x+16, y+4)
	board.Seed(x+11, y+5)
	board.Seed(x+16, y+5)
	board.Seed(x+12, y+6)
	board.Seed(x+15, y+6)
	board.Seed(x+13, y+7)
	board.Seed(x+14, y+7)
	board.Seed(x+13, y+8)
	board.Seed(x+14, y+8)
	board.Seed(x+25, y+1)
	board.Seed(x+25, y+2)
	board.Seed(x+25, y+3)
	board.Seed(x+26, y+1)
	board.Seed(x+27, y+2)
	board.Seed(x+26, y+3)
	board.Seed(x+35, y+2)
	board.Seed(x+36, y+2)
	board.Seed(x+35, y+3)
	board.Seed(x+36, y+3)
}
