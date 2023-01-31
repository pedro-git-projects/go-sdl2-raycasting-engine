package main

import (
	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

func render(r *sdl.Renderer, g *game.Game, p *player.Player) {
	r.SetDrawColor(0, 0, 0, 255)
	r.Clear()
	g.RenderMap(r)
	p.Render(r)
	r.Present()
}
