package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	g, p := setup()
	w, r, buff, err := initializeWindow(g)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()
	defer w.Destroy()

	running := true
	for running {
		processInput(&running, p)
		update(g, p)
		render(r, g, p, buff)
	}
}
