package main

import (
	"math"
	"unsafe"

	"github.com/pedro-git-projects/go-raycasting/cmd/window"
	"github.com/veandco/go-sdl2/sdl"
)

// render renders game objects into the window
func (app *App) render() {
	app.renderer.SetDrawColor(0, 0, 0, 255)
	app.renderer.Clear()

	app.generate3DProjection()

	app.renderColorBuffer()
	app.colorBuffer.Clear(0xFF000000)

	app.renderMap()
	app.renderRays()
	app.renderPlayer()
	app.renderer.Present()
}

// renderPlayer renders a line to represent the current player position in the minimap
func (app *App) renderPlayer() {
	app.renderer.SetDrawColor(255, 255, 255, 255)
	p := sdl.Rect{
		X: int32(app.player.X() * window.MinimapScaling),
		Y: int32(app.player.Y() * window.MinimapScaling),
		W: int32(app.player.Width() * window.MinimapScaling),
		H: int32(app.player.Height() * window.MinimapScaling),
	}
	app.renderer.FillRect(&p)

	length := 30.0
	app.renderer.DrawLine(
		int32(window.MinimapScaling*app.player.X()),
		int32(window.MinimapScaling*app.player.Y()),
		int32((window.MinimapScaling)*app.player.X()+math.Cos(app.player.RotationAngle())*length),
		int32((window.MinimapScaling)*app.player.Y()+math.Sin(app.player.RotationAngle())*length),
	)
}

// renderMap renders the game minimap as well as a rectangle behind it
// so as not to let the walls be seen through the tile gaps
func (app *App) renderMap() {
	app.renderer.SetDrawColor(0, 0, 0, 255)
	app.renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(math.Floor(window.MinimapScaling * (float64(window.NumCols) * float64(window.TileSize)))),
		H: int32(math.Floor(window.MinimapScaling * (float64(window.NumRows) * float64(window.TileSize)))),
	})

	for i := int32(0); i < window.NumRows; i++ {
		for j := int32(0); j < window.NumCols; j++ {
			xTile := j * window.TileSize
			yTile := i * window.TileSize

			var tileColor uint8 = 0
			if app.game.GameMap()[i][j] != 0 {
				tileColor = 255
			}

			app.renderer.SetDrawColor(tileColor, tileColor, tileColor, 255)
			mapTile := sdl.Rect{
				X: int32(float64(xTile) * window.MinimapScaling),
				Y: int32(float64(yTile) * window.MinimapScaling),
				W: int32(math.Floor(float64(window.TileSize) * window.MinimapScaling)),
				H: int32(math.Floor(float64(window.TileSize) * window.MinimapScaling)),
			}
			app.renderer.FillRect(&mapTile)
		}
	}
}

// renderRays draws as many lines as the number of rays
// they start at the player position and stop at the first collision detection coodrinate
// in that line
func (app *App) renderRays() {
	app.renderer.SetDrawColor(255, 0, 0, 255)
	for i := 0; i < int(window.NumRays); i++ {
		app.renderer.DrawLine(
			int32(window.MinimapScaling*app.player.X()),
			int32(window.MinimapScaling*app.player.Y()),
			int32(window.MinimapScaling*app.game.Rays()[i].XCollision()),
			int32(window.MinimapScaling*app.game.Rays()[i].YCollision()),
		)
	}
}

// renderColorBuffer  unsafe calls to sdl Texture Update and Copy via cgo
func (app *App) renderColorBuffer() {
	app.colorBuffer.Texture().Update(
		nil,
		unsafe.Pointer(&app.colorBuffer.Color()[0]),
		int(uint32(window.Width)*uint32(unsafe.Sizeof(uint32(0)))),
	)
	app.renderer.Copy(app.colorBuffer.Texture(), nil, nil)
}

// Generate3DProjection accounts for the distortion by calculating the perpendicular distance between the player
// and the collided pixel then calculates the projected top and bottom pixel positions before drawing a line there
// it also uses the fact that the collision was vertical or horizontal to cast "shadows"
func (app *App) generate3DProjection() {
	for x := int32(0); x < window.NumRays; x++ {
		perpDist := app.game.Rays()[x].Distance() * math.Cos(app.game.Rays()[x].Angle()-app.player.RotationAngle())
		projWallHeight := (float64(window.TileSize) / perpDist) * window.DistanceProjPlane
		wallSegmentHeight := int32(projWallHeight)

		var topWallPixel int32 = (window.Height / 2) - (wallSegmentHeight / 2)
		if topWallPixel < 0 {
			topWallPixel = 0
		}

		var bottomWallPixel int32 = (window.Height / 2) + (wallSegmentHeight / 2)
		if bottomWallPixel > window.Height {
			bottomWallPixel = window.Height
		}

		// ceiling
		for y := int32(0); y < topWallPixel; y++ {
			app.colorBuffer.Color()[(window.Width*y)+x] = 0xFF444444
		}

		// wall
		for y := topWallPixel; y < bottomWallPixel; y++ {
			if app.game.Rays()[x].IsVerticalCollision() {
				app.colorBuffer.Color()[(window.Width*y)+x] = 0xFFFFFFFF
			} else {
				app.colorBuffer.Color()[(window.Width*y)+x] = 0xFFCCCCCC
			}
		}

		// floor
		for y := bottomWallPixel; y < window.Height; y++ {
			app.colorBuffer.Color()[(window.Width*y)+x] = 0xFF777777
		}
	}
}
