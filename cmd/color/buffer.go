package color

import (
	"unsafe"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/veandco/go-sdl2/sdl"
)

type Buffer struct {
	color   []uint32
	texture *sdl.Texture
}

func New(r *sdl.Renderer, g *game.Game) *Buffer {
	// maybe come back and remove multiplication
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
