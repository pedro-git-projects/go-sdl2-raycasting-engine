package game

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	windowWidth    int32
	windowHeight   int32
	tileSize       int32
	rows           int32
	cols           int32
	minimapScale   float64
	fovAngle       float64
	rays           int32
	fps            uint64
	frameTime      uint64
	ticksLastFrame uint64
	gameMap        [][]int32
}

func Default() *Game {
	framesPerSecond := uint64(30)
	rows := int32(13)
	cols := int32(20)
	tileSiz := int32(64)
	g := &Game{
		windowWidth:    cols * tileSiz,
		windowHeight:   rows * tileSiz,
		tileSize:       tileSiz,
		rows:           rows,
		cols:           cols,
		minimapScale:   0.3,
		rays:           cols * tileSiz,
		fovAngle:       (60 * (math.Pi / 180)),
		gameMap:        initializeGameMap(),
		fps:            framesPerSecond,
		frameTime:      1000 / framesPerSecond,
		ticksLastFrame: 0,
	}
	return g
}

func initializeGameMap() [][]int32 {
	m := [][]int32{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	return m
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

func (g *Game) RenderMap(r *sdl.Renderer) {
	for i := int32(0); i < g.rows; i++ {
		for j := int32(0); j < g.cols; j++ {

			tileX := j * g.tileSize
			tileY := i * g.tileSize

			var tileColor uint8
			if g.gameMap[i][j] != 0 {
				tileColor = 255
			} else {
				tileColor = 0
			}

			r.SetDrawColor(tileColor, tileColor, tileColor, 255)
			mapTileRect := sdl.Rect{
				X: int32(float64(tileX) * g.minimapScale),
				Y: int32(float64(tileY) * g.minimapScale),
				W: int32(float64(g.tileSize) * g.minimapScale),
				H: int32(float64(g.tileSize) * g.minimapScale),
			}
			r.FillRect(&mapTileRect)
		}
	}
}

func (g *Game) IsSolidCoordinate(x, y float64) bool {
	if x < 0 || x >= float64(g.cols)*float64(g.tileSize) || y < 0 || y >= float64(g.rows)*float64(g.tileSize) {
		return true
	}

	indX := math.Floor(x / float64(g.tileSize))
	indY := math.Floor(y / float64(g.tileSize))
	if g.gameMap[int(indY)][int(indX)] != 0 {
		return true
	}

	return false
}
