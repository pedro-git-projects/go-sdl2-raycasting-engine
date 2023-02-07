package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// update gets the durations since SDL started and compensates for the
// difference between the actual framerate and the frametime, or time between frames// by wasting time away when the game is running too fast
func (app *App) update() {
	timeToWait := app.timer.FrameTime() - (sdl.GetTicks64() - app.game.TicksLastFrame())
	if timeToWait > 0 && timeToWait <= app.timer.FrameTime() {
		time.Sleep(time.Duration(timeToWait) * time.Millisecond)
	}
	deltaTime := (float64(sdl.GetTicks64() - app.game.TicksLastFrame())) / 1000.0
	app.game.SetTicksLastFrame(sdl.GetTicks64())
	app.player.Move(deltaTime, app.game)
	app.CastRays()
}
