package game

import (
	"math"

	"github.com/pedro-git-projects/go-raycasting/cmd/ray"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	windowWidth       int32
	windowHeight      int32
	tileSize          int32
	rows              int32
	cols              int32
	minimapScale      float64
	fovAngle          float64
	numRays           int32
	fps               uint64
	frameTime         uint64
	ticksLastFrame    uint64
	gameMap           [][]int32
	distanceProjPlane float64
	rays              []ray.Ray
}

func Default() *Game {
	framesPerSecond := uint64(30)
	rows := int32(13)
	cols := int32(20)
	tileSiz := int32(64)
	width := cols * tileSiz
	height := rows * tileSiz
	fov := 60 * (math.Pi / 180)
	dist := ((float64(width) / 2) / (math.Tan(fov / 2)))
	numRays := cols * tileSiz
	rays := make([]ray.Ray, numRays)
	g := &Game{
		windowWidth:       width,
		windowHeight:      height,
		tileSize:          tileSiz,
		rows:              rows,
		cols:              cols,
		minimapScale:      0.3,
		numRays:           width,
		fovAngle:          fov,
		gameMap:           initializeGameMap(),
		fps:               framesPerSecond,
		frameTime:         1000 / framesPerSecond,
		ticksLastFrame:    0,
		distanceProjPlane: dist,
		rays:              rays,
	}
	return g
}

func initializeGameMap() [][]int32 {
	m := [][]int32{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 2, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 5},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 5},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 5},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5},
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

func (g Game) FOV() float64 {
	return g.fovAngle
}

func (g Game) NumRays() int32 {
	return g.numRays
}

func (g Game) Rays() []ray.Ray {
	return g.rays
}

func (g Game) TileSize() int32 {
	return g.tileSize
}

func (g Game) Cols() int32 {
	return g.cols
}

func (g Game) Rows() int32 {
	return g.rows
}

func (g Game) GameMap() [][]int32 {
	return g.gameMap
}

func (g Game) DistanceProjectionPlane() float64 {
	return g.distanceProjPlane
}

func (g Game) MinimapScale() float64 {
	return g.minimapScale
}

func (g *Game) RenderMap(r *sdl.Renderer) {
	for i := int32(0); i < g.rows; i++ {
		for j := int32(0); j < g.cols; j++ {

			tileX := j * g.tileSize
			tileY := i * g.tileSize

			var tileColor uint8 = 0
			if g.gameMap[i][j] != 0 {
				tileColor = 255
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
	if x < 0 || x >= float64(g.windowWidth) ||
		y < 0 || y > float64(g.windowHeight) {
		return true
	}

	indX := int(math.Floor(x / float64(g.tileSize)))
	indY := int(math.Floor(y / float64(g.tileSize)))
	return g.gameMap[indY][indX] != 0
}
