package main

import (
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

func processInput(running *bool, player *player.Player) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch i := event.(type) {
		case *sdl.QuitEvent:
			*running = false
			break
		case *sdl.KeyboardEvent:
			if i.Keysym.Sym == sdl.K_ESCAPE {
				*running = false
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_UP {
				player.SetWalkDirection("foward")
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_DOWN {
				player.SetWalkDirection("down")
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_RIGHT {
				player.SetTurnDirection("right")
			}
			if i.Type == sdl.KEYDOWN && i.Keysym.Sym == sdl.K_LEFT {
				player.SetTurnDirection("left")
			}

			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_UP {
				player.SetWalkDirection("neutral")
			}
			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_DOWN {
				player.SetWalkDirection("neutral")
			}
			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_RIGHT {
				player.SetTurnDirection("neutral")
			}
			if i.Type == sdl.KEYUP && i.Keysym.Sym == sdl.K_LEFT {
				player.SetTurnDirection("neutral")
			}
			break
		}
	}
}
