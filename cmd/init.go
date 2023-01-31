package main

import (
	"errors"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

func initializeWindow(g *game.Game) (*sdl.Window, *sdl.Renderer, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, nil, errors.New("Failed to initialze SDL")
	}

	window, err := sdl.CreateWindow(
		"raycasting",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		g.WindowWidth(),
		g.WindowHeight(),
		sdl.WINDOW_BORDERLESS,
	)
	if err != nil {
		return nil, nil, err
	}

	rendeder, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		return nil, nil, err
	}

	err = rendeder.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		return nil, nil, err
	}

	return window, rendeder, nil
}

func setup() (*game.Game, *player.Player) {
	g := game.Default()
	p := player.New(float64(g.WindowWidth()/2), float64(g.WindowHeight()/2))
	return g, p
}
