// timekeepr stores all datatypes, constants, variables and functions relating to
// framerate, frametime and update pacing
package timekeeper

import (
	"github.com/veandco/go-sdl2/sdl"
)

// default values
const (
	fps uint64 = 60
)

var (
	frameTimeTarget = 1000 / fps
)

// TimeKeeper represents all application data that is related to framerate, frametime and update pace
type TimeKeeper struct {
	ticksLastFrame  uint64 // number of milliseconds ellapsed since last frame
	deltaTime       float64
	fps             uint64
	frameTimeTarget uint64 // number of milliseconds each frame should take
	waitTime        uint64
}

func Default() *TimeKeeper {
	t := &TimeKeeper{
		ticksLastFrame:  0,
		fps:             fps,
		frameTimeTarget: frameTimeTarget,
	}
	return t
}

// TicksLastFrame returns the number of milliseconds ellapsed since last frame
func (t TimeKeeper) TicksLastFrame() uint64 {
	return t.ticksLastFrame
}

// SetTicks sets the number of milliseconds ellapsed since last frame
func (t *TimeKeeper) SetTicks(ticks uint64) {
	t.ticksLastFrame = ticks
}

// SetDelta sets the current delta time
func (t *TimeKeeper) SetDelta(newDelta float64) {
	t.deltaTime = newDelta
}

// Accessor for the current delta time
func (t TimeKeeper) DeltaTime() float64 {
	return t.deltaTime
}

// CalculateDelta calcualtes and sets the application delta time
func (t *TimeKeeper) CalculateDelta() {
	d := float64((sdl.GetTicks64() - t.TicksLastFrame())) / 1000.0
	t.deltaTime = d
}

// CalculateWaitTime calculates and sets how much time the application must stop to
// respect the frame time
func (t *TimeKeeper) CalculateWaitTime() {
	w := t.frameTimeTarget - (sdl.GetTicks64() - t.ticksLastFrame)
	t.waitTime = w
}

/* Accessors */

func (t TimeKeeper) FPS() uint64 {
	return t.fps
}

func (t TimeKeeper) FrameTime() uint64 {
	return t.frameTimeTarget
}

func (t TimeKeeper) WaitTime() uint64 {
	return t.waitTime
}
