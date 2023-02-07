package game

import (
	"math"

	"github.com/pedro-git-projects/go-raycasting/cmd/ray"
	"github.com/pedro-git-projects/go-raycasting/cmd/window"
)

type Game struct {
	tileSize          int32
	ticksLastFrame    uint64
	gameMap           [][]int32
	distanceProjPlane float64
	rays              []ray.Ray
}

// Default creates a game object with its fields populated by the
// default constants and variables
func Default() *Game {
	rays := make([]ray.Ray, window.NumRays)
	g := &Game{
		tileSize:          window.TileSize,
		gameMap:           initializeGameMap(),
		ticksLastFrame:    0,
		distanceProjPlane: window.DistanceProjPlane,
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

func (g Game) TicksLastFrame() uint64 {
	return g.ticksLastFrame
}

func (g Game) Rays() []ray.Ray {
	return g.rays
}

func (g Game) TileSize() int32 {
	return window.TileSize
}

func (g Game) GameMap() [][]int32 {
	return g.gameMap
}

func (g Game) DistanceProjectionPlane() float64 {
	return g.distanceProjPlane
}

// IsSolidCoordinate tests if a point x,y is solid
// that is, has collision. It will return true if the
// unrwaped value is not zero or if the point is out of bounds
func (g *Game) IsSolidCoordinate(x, y float64) bool {
	if x < 0 || x >= float64(window.Width) ||
		y < 0 || y > float64(window.Height) {
		return true
	}

	indX := int(math.Floor(x / float64(window.TileSize)))
	indY := int(math.Floor(y / float64(window.TileSize)))
	return g.gameMap[indY][indX] != 0
}

// Mutator
func (g *Game) SetTicksLastFrame(t uint64) {
	g.ticksLastFrame = t
}
