package game

type Game struct {
	windowWidth    int32
	windowHeight   int32
	fps            uint64
	frameTime      uint64
	ticksLastFrame uint64
}

func Default() *Game {
	framesPerSecond := uint64(30)
	g := &Game{
		windowWidth:    800,
		windowHeight:   600,
		fps:            framesPerSecond,
		frameTime:      1000 / framesPerSecond,
		ticksLastFrame: 0,
	}
	return g
}

func (g Game) WindowWidth() int32 {
	return g.windowWidth
}

func (g Game) WindowHeight() int32 {
	return g.windowHeight
}

func (g Game) FrameTime() uint64 {
	return g.frameTime
}

func (g Game) TicksLastFrame() uint64 {
	return g.ticksLastFrame
}

func (g *Game) SetTicksLastFrame(t uint64) {
	g.ticksLastFrame = t
}
