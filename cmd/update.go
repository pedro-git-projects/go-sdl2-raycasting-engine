package main

import (
	"time"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

// update gets the durations since SDL started and compensates for the
// difference between the actual framerate and the frametime, or time between frames// by wasting time away when the game is running too fast
func update(g *game.Game, p *player.Player) {
	timeToWait := g.FrameTime() - (sdl.GetTicks64() - g.TicksLastFrame())
	if timeToWait > 0 && timeToWait <= g.FrameTime() {
		time.Sleep(time.Duration(timeToWait) * time.Millisecond)
	}
	deltaTime := (float64(sdl.GetTicks64() - g.TicksLastFrame())) / 1000.0
	g.SetTicksLastFrame(sdl.GetTicks64())
	p.Move(deltaTime, g)
}
