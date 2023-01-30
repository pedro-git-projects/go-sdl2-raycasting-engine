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

func setup() (*player.Player, *game.Game) {
	p := player.New(0, 0)
	g := game.Default()
	return p, g
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

// TODO: fix get ticks passed
func getTicksPassed(a, b uint64) bool {
	r := ((b) - (a)) <= 0
	return r
}

// Create frame independent movement
func update(p *player.Player, g *game.Game) {
	//	for !getTicksPassed(sdl.GetTicks64(), g.TicksLastFrame()+g.FrameTime()) {
	//	}

	//delta := (sdl.GetTicks64() - g.TicksLastFrame()) / 1000.0
	g.SetTicksLastFrame(sdl.GetTicks64())

	p.IncX(1)
	p.IncY(1)
}

func main() {
	p, g := setup()
	w, r, err := initializeWindow(g)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()
	defer w.Destroy()

	running := true
	for running {
		processInput(&running)
		update(p, g)
		render(r, p)
	}
}
