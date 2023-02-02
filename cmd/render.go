package main

import (
	"github.com/pedro-git-projects/go-raycasting/cmd/color"
	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/pedro-git-projects/go-raycasting/cmd/ray/cast"
	"github.com/veandco/go-sdl2/sdl"
)

func render(r *sdl.Renderer, g *game.Game, p *player.Player, buf *color.Buffer) {
	r.SetDrawColor(0, 0, 0, 255)
	r.Clear()

	buf.Generate3DProjection(g, p)

	buf.Render(g, r)
	buf.ClearColorBuffer(0xFF000000, g)

	g.RenderMap(r)
	cast.RenderRays(r, g, p)
	p.Render(r)
	r.Present()
}
