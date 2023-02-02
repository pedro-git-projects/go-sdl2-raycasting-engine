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
			break
		case *sdl.KeyboardEvent:
			if i.Keysym.Sym == sdl.K_ESCAPE {
				app.SetRunning(false)
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_UP {
				app.player.SetWalkDirection("foward")
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_DOWN {
				app.player.SetWalkDirection("backward")
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_RIGHT {
				app.player.SetTurnDirection("right")
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_LEFT {
				app.player.SetTurnDirection("left")
			}

			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_UP {
				app.player.SetWalkDirection("neutral")
			}
			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_DOWN {
				app.player.SetWalkDirection("neutral")
			}
			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_RIGHT {
				app.player.SetTurnDirection("neutral")
			}
			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_LEFT {
				app.player.SetTurnDirection("neutral")
			}
			break
		}
	}
}
