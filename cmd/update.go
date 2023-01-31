package main

import (
	"time"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/veandco/go-sdl2/sdl"
)

// update gets the durations since SDL started and compensates for the
// difference between the actual framerate and the frametime, or time between frames// by wasting time away when the game is running too fast
func update(g *game.Game) {
	g.SetTicksLastFrame(sdl.GetTicks64())

	delta := (sdl.GetTicks64() - g.TicksLastFrame())
	if delta < g.FrameTime() {
		sleep := (g.FrameTime() - delta) * uint64(time.Millisecond)
		time.Sleep(time.Duration(sleep))
	}
}
