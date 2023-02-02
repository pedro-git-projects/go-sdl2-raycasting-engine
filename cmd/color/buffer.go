package color

import (
	"math"
	"unsafe"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

type Buffer struct {
	color   []uint32
	texture *sdl.Texture
}

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

func (b *Buffer) Render(g *game.Game, r *sdl.Renderer) {
	b.texture.Update(
		nil,
		unsafe.Pointer(&b.color[0]),
		int(uint32(g.WindowWidth())*uint32(unsafe.Sizeof(uint32(0)))),
	)
	r.Copy(b.texture, nil, nil)
}

func (b *Buffer) ClearColorBuffer(color uint32, g *game.Game) {
	for i := 0; i < int(g.WindowWidth()*g.WindowHeight()); i++ {
		b.color[i] = color
	}
}

func (b *Buffer) Generate3DProjection(g *game.Game, p *player.Player) {
	for x := 0; x < int(g.NumRays()); x++ {
		perpendicularDist := g.Rays()[x].Distance() * math.Cos(g.Rays()[x].Angle()-p.RotationAngle())
		projectedWallHeight := float64(g.TileSize()) / perpendicularDist * g.DistanceProjectionPlane()
		topWallPixel := int(float64(g.WindowHeight())/2 - (projectedWallHeight / 2))
		if topWallPixel < 0 {
			topWallPixel = 0
		}

		bottomWallPixel := int(float64(g.WindowHeight())/2 + (projectedWallHeight / 2))
		if bottomWallPixel > int(g.WindowHeight()) {
			bottomWallPixel = int(g.WindowHeight())
		}

		for y := topWallPixel; y < bottomWallPixel; y++ {
			if g.Rays()[x].IsVerticalCollision() {
				b.color[(int(g.WindowWidth())*y)+x] = 0xFFFFFFFF
			} else {
				b.color[(int(g.WindowWidth())*y)+x] = 0xFFCCCCCC
			}

		}

	}
}
