package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// processInput processes player input using sdl's PollEvent function
func (app *App) processInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch i := event.(type) {
		case *sdl.QuitEvent:
			app.SetRunning(false)
		case *sdl.KeyboardEvent:
			key := i.Keysym.Sym
			if i.Type == sdl.KEYDOWN {
				switch key {
				case sdl.K_ESCAPE:
					app.SetRunning(false)
				case sdl.K_UP:
					app.player.SetWalkDirection("foward")
				case sdl.K_DOWN:
					app.player.SetWalkDirection("backward")
				case sdl.K_RIGHT:
					app.player.SetTurnDirection("right")
				case sdl.K_LEFT:
					app.player.SetTurnDirection("left")
				}
			}
			if i.Type == sdl.KEYUP {
				switch key {
				case sdl.K_UP, sdl.K_DOWN:
					app.player.SetWalkDirection("neutral")
				case sdl.K_RIGHT, sdl.K_LEFT:
					app.player.SetTurnDirection("neutral")
				}
			}
		}
	}
}
