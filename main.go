package main

import (
	"time"

	"github.com/callegarimattia/go-life/model"
	"github.com/callegarimattia/go-life/view"
)

func main() {
	b := model.NewBoard(20, 20)

	for {
		view.Render(b)
		b = b.Tick()

		<-time.After(500 * time.Millisecond)
	}
}
