package game

import (
	"math"

	"github.com/pedro-git-projects/go-raycasting/cmd/ray"
	"github.com/veandco/go-sdl2/sdl"
)

// default values
const (
	MinimapScaling float64 = 0.2
	tileSize       int32   = 64
	numRows        int32   = 13
	numCols        int32   = 20
	fov            float64 = 60 * (math.Pi / 180)
	fps            uint64  = 30
)

var (
	width             int32   = numCols * tileSize
	height            int32   = numRows * tileSize
	distanceProjPlane float64 = ((float64(width) / 2) / math.Tan(fov/2))
	frameTime         uint64  = 1000 / fps
	numRays                   = width
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

// Default creates a game object with its fields populated by the
// default constants and variables
func Default() *Game {
	rays := make([]ray.Ray, numRays)
	g := &Game{
		windowWidth:       width,
		windowHeight:      height,
		tileSize:          tileSize,
		rows:              numRows,
		cols:              numCols,
		minimapScale:      MinimapScaling,
		numRays:           width,
		fovAngle:          fov,
		gameMap:           initializeGameMap(),
		fps:               fps,
		frameTime:         frameTime,
		ticksLastFrame:    0,
		distanceProjPlane: distanceProjPlane,
		rays:              rays,
	}
	return g
}

// initializeGameMap is a constructor helper for the gameMap field
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

/* Accessors */

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

// RenderMap renders the game minimap as well as a rectangle behind it
// so as not to let the walls be seen through the tile gaps
func (g *Game) RenderMap(r *sdl.Renderer) {
	r.SetDrawColor(0, 0, 0, 255)
	r.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(g.MinimapScale() * (float64(g.cols) * float64(g.TileSize()))),
		H: int32(g.MinimapScale() * (float64(g.Rows()) * float64(g.TileSize()))),
	})

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

// IsSolidCoordinate tests if a point x,y is solid
// that is, has collision. It will return true if the
// unrwaped value is not zero or if the point is out of bounds
func (g *Game) IsSolidCoordinate(x, y float64) bool {
	if x < 0 || x >= float64(g.windowWidth) ||
		y < 0 || y > float64(g.windowHeight) {
		return true
	}

	indX := int(math.Floor(x / float64(g.tileSize)))
	indY := int(math.Floor(y / float64(g.tileSize)))
	return g.gameMap[indY][indX] != 0
}

// Mutator
func (g *Game) SetTicksLastFrame(t uint64) {
	g.ticksLastFrame = t
}
