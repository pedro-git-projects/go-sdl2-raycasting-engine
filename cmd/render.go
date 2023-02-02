package main

import (
	"github.com/pedro-git-projects/go-raycasting/cmd/ray/cast"
)

// render renders game objects into the window
func (app *App) render() {
	app.renderer.SetDrawColor(0, 0, 0, 255)
	app.renderer.Clear()

	app.colorBuffer.Generate3DProjection(app.game, app.player)

	app.colorBuffer.Render(app.game, app.renderer)
	app.colorBuffer.ClearColorBuffer(0xFF000000, app.game)

	app.game.RenderMap(app.renderer)
	cast.RenderRays(app.renderer, app.game, app.player)
	app.player.Render(app.renderer)
	app.renderer.Present()
}
