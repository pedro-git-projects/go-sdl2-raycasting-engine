package main

import (
	"errors"

	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

func initializeWindow() (*sdl.Window, *sdl.Renderer, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, nil, errors.New("Failed to initialze SDL")
	}

	window, err := sdl.CreateWindow(
		"raycasting",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		800,
		600,
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

func setup() *player.Player {
	p := player.New(0, 0)
	return p
}

func processInput(running *bool) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch i := event.(type) {
		case *sdl.QuitEvent:
			*running = false
			break
		case *sdl.KeyboardEvent:
			if i.Keysym.Sym == sdl.K_ESCAPE {
				*running = false
				break
			}
		}
	}
}

func render(r *sdl.Renderer, p *player.Player) {
	r.SetDrawColor(0, 0, 0, 255)
	r.Clear()

	r.SetDrawColor(255, 255, 0, 255)
	rect := sdl.Rect{int32(p.X()), int32(p.Y()), 20, 20}
	r.FillRect(&rect)

	r.Present()
}

func update(p *player.Player) {
	p.IncX(1)
	p.IncY(1)
}

func main() {
	w, r, err := initializeWindow()
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()
	defer w.Destroy()

	p := setup()

	running := true
	for running {
		processInput(&running)
		update(p)
		render(r, p)
	}
}
