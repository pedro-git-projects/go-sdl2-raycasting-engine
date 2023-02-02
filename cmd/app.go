package main

import (
	"errors"

	"github.com/pedro-git-projects/go-raycasting/cmd/color"
	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/veandco/go-sdl2/sdl"
)

// the App struct encaplsulates all parts of the application whilst also serving al
// a mean of achieving dependency injection by making what used to be functions
// with a variety of pointers passed now app pointer recievers without parameters
type App struct {
	game        *game.Game
	window      *sdl.Window
	renderer    *sdl.Renderer
	colorBuffer *color.Buffer
	player      *player.Player
	isRunning   bool
}

// newApp returns a pointer to an App populated with a new game and player,
// other fields must be populated by calling initializeWindow
func newApp() *App {
	g := game.Default()
	p := player.New(float64(g.WindowWidth())/2, float64(g.WindowHeight())/2)
	a := &App{
		game:   g,
		player: p,
	}
	return a
}

// initialize initializes SDL and populates the missing fileds in the current App instance
func (app *App) initialize() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return errors.New("Failed to initialze SDL")
	}

	window, err := sdl.CreateWindow(
		"raycasting",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		app.game.WindowWidth(),
		app.game.WindowHeight(),
		sdl.WINDOW_BORDERLESS,
	)
	if err != nil {
		return err
	}
	app.window = window

	rendeder, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		return err
	}
	app.renderer = rendeder

	err = rendeder.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		return err
	}

	colorBuffer := color.New(rendeder, app.game)
	app.colorBuffer = colorBuffer

	app.SetRunning(true)

	return nil
}

// destructor frees system resources and destroys the window
// it should be called in a defer statement
func (app *App) destructor() {
	sdl.Quit()
	app.window.Destroy()
}

// IsRunning is an accessor for the running field
func (app *App) IsRunning() bool {
	return app.isRunning
}

// SetRunning is a setter for the isRunning filed
func (app *App) SetRunning(isRunning bool) {
	app.isRunning = isRunning
}
