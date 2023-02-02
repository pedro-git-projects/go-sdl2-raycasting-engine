package color

import (
	"math"
	"unsafe"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

// Buffer represents a color buffer
type Buffer struct {
	color   []uint32
	texture *sdl.Texture
}

// Creates a new Buffer with a streaming texture of the width and height of the window
// and populates the color field with a slice that is wxh by wxh
// Note that the slice is of type uint32 because that is the type assigned to colors in SDL
func New(r *sdl.Renderer, g *game.Game) *Buffer {
	color := make([]uint32, g.WindowWidth()*g.WindowHeight(), g.WindowHeight()*g.WindowWidth())
	texture, _ := r.CreateTexture(
		sdl.PIXELFORMAT_ARGB8888,
		sdl.TEXTUREACCESS_STREAMING,
		g.WindowWidth(),
		g.WindowHeight(),
	)
	b := &Buffer{
		color:   color,
		texture: texture,
	}
	return b
}

// Render makes unsafe calls to sdl Texture Update and Copy via cgo
func (b *Buffer) Render(g *game.Game, r *sdl.Renderer) {
	b.texture.Update(
		nil,
		unsafe.Pointer(&b.color[0]),
		int(uint32(g.WindowWidth())*uint32(unsafe.Sizeof(uint32(0)))),
	)
	r.Copy(b.texture, nil, nil)
}

// Clear draws over the whole screen with an specified color
func (b *Buffer) Clear(color uint32, g *game.Game) {
	for i := 0; i < int(g.WindowWidth()*g.WindowHeight()); i++ {
		b.color[i] = color
	}
}

// Generate3DProjection accounts for the distortion by calculating the perpendicular distance between the player
// and the collided pixel then calculates the projected top and bottom pixel positions before drawing a line there
// it also uses the fact that the collision was vertical or horizontal to cast "shadows"
func (b *Buffer) Generate3DProjection(g *game.Game, p *player.Player) {
	for x := int32(0); x < g.NumRays(); x++ {

		perpDist := g.Rays()[x].Distance() * math.Cos(g.Rays()[x].Angle()-p.RotationAngle())
		projWallHeight := (float64(g.TileSize()) / perpDist) * g.DistanceProjectionPlane()
		wallSegmentHeight := int32(projWallHeight)

		var topWallPixel int32 = (g.WindowHeight() / 2) - (wallSegmentHeight / 2)
		if topWallPixel < 0 {
			topWallPixel = 0
		}

		var bottomWallPixel int32 = (g.WindowHeight() / 2) + (wallSegmentHeight / 2)
		if bottomWallPixel > g.WindowHeight() {
			bottomWallPixel = g.WindowHeight()
		}

		// ceiling
		for y := int32(0); y < topWallPixel; y++ {
			b.color[(g.WindowWidth()*y)+x] = 0xFF444444
		}

		// wall
		for y := topWallPixel; y < bottomWallPixel; y++ {
			if g.Rays()[x].IsVerticalCollision() {
				b.color[(g.WindowWidth()*y)+x] = 0xFFFFFFFF
			} else {
				b.color[(g.WindowWidth()*y)+x] = 0xFFCCCCCC
			}
		}

		// floor
		for y := bottomWallPixel; y < g.WindowHeight(); y++ {
			b.color[(g.WindowWidth()*y)+x] = 0xFF777777
		}
	}
}
