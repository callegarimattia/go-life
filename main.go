package main

import (
	"os"
	"time"

	"github.com/callegarimattia/go-life/model"
	"github.com/callegarimattia/go-life/shapes"
	"github.com/callegarimattia/go-life/view"
)

func main() {
	b := model.NewBoard(64, 16)

	shapes.SeedDiehard(b, 10, 5)

	ticking(200*time.Millisecond, b)
}

func ticking(interval time.Duration, b model.Board) {
	for {
		view.Render(os.Stdout, b)
		b = b.Tick()

		<-time.After(interval)
	}
}
