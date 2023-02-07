package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// update gets the durations since SDL started and compensates for the
// difference between the actual framerate and the frametime, or time between frames
// by wasting time away when the game is running too fast
func (app *App) update() {
	app.timer.CalculateWaitTime()

	if app.timer.WaitTime() > 0 && app.timer.WaitTime() <= app.timer.FrameTime() {
		sdl.Delay(uint32(app.timer.WaitTime())) //	time.Sleep(time.Duration(app.timer.WaitTime()) * time.Millisecond)
	}

	app.timer.CalculateDelta()
	app.timer.SetTicks(sdl.GetTicks64())
	app.player.Move(app.timer.DeltaTime(), app.game)
	app.CastRays()
}
