package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	g, _ := setup()
	w, r, err := initializeWindow(g)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()
	defer w.Destroy()

	running := true
	for running {
		processInput(&running)
		update(g)
		render(r, g)
	}
}
