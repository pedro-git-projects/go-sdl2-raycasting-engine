package color

import (
	"github.com/pedro-git-projects/go-raycasting/cmd/window"
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
func New(r *sdl.Renderer) *Buffer {
	color := make([]uint32, window.Width*window.Height, window.Height*window.Width)
	texture, _ := r.CreateTexture(
		sdl.PIXELFORMAT_ARGB8888,
		sdl.TEXTUREACCESS_STREAMING,
		window.Width,
		window.Height,
	)
	b := &Buffer{
		color:   color,
		texture: texture,
	}
	return b
}

// Clear draws over the whole screen with an specified color
func (b *Buffer) Clear(color uint32) {
	for i := 0; i < int(window.Width*window.Height); i++ {
		b.color[i] = color
	}
}

/* Accessors */

func (b Buffer) Texture() *sdl.Texture {
	return b.texture
}

func (b Buffer) Color() []uint32 {
	return b.color
}
