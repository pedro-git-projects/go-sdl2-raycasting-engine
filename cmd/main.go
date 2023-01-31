package main

import (
	"errors"
	"time"

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

// update gets the durations since SDL started and compensates for the
// difference between the actual framerate and the frametime, or time between frames// by wasting time away when the game is running too fast
func update(p *player.Player, g *game.Game) {
	g.SetTicksLastFrame(sdl.GetTicks64())

	delta := (sdl.GetTicks64() - g.TicksLastFrame())
	if delta < g.FrameTime() {
		sleep := (g.FrameTime() - delta) * uint64(time.Millisecond)
		time.Sleep(time.Duration(sleep))
	}

	p.IncX(20)
	p.IncY(20)
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
