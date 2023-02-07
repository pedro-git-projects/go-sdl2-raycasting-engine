package timekeeper

import "github.com/veandco/go-sdl2/sdl"

// default values
const (
	fps uint64 = 30
)

var (
	frameTimeTarget = 1000 / fps
)

type TimeKeeper struct {
	ticksLastFrame  uint64 // number of milliseconds ellapsed since last frame
	deltaTime       uint64
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
func (t *TimeKeeper) SetDelta(newDelta uint64) {
	t.deltaTime = newDelta
}

// Accessor for the current delta time
func (t TimeKeeper) DeltaTime() uint64 {
	return t.deltaTime
}

func (t *TimeKeeper) CalculateDelta() {
	d := ((sdl.GetTicks64() - t.TicksLastFrame()) / 1000)
	t.deltaTime = d
}

func (t TimeKeeper) FPS() uint64 {
	return t.fps
}

func (t TimeKeeper) FrameTime() uint64 {
	return t.frameTimeTarget
}

func (t *TimeKeeper) CalculateWaitTime() {
	w := t.frameTimeTarget - (sdl.GetTicks64() - t.ticksLastFrame)
	t.waitTime = w
}

func (t TimeKeeper) WaitTime() uint64 {
	return t.waitTime
}
