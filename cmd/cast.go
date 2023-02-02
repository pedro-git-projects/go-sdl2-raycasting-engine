package main

import (
	"math"

	"github.com/pedro-git-projects/go-raycasting/cmd/ray"
	"github.com/pedro-git-projects/go-raycasting/cmd/utils"
)

func (app *App) CastRay(angle float64, rayId int) {
	app.game.Rays()[rayId] = *ray.New(angle) // constructor sanitizes angle
	ray := app.game.Rays()[rayId]

	var horzXWallCollision float64
	var horzYWallCollision float64
	var horzWallContent int32
	foundHorzCollision := false

	// closest y-coordinate to the horizontal grid intersection
	yIntersection := math.Floor(app.player.Y()/float64(app.game.TileSize())) * float64(app.game.TileSize())
	if ray.IsFacingDown() {
		yIntersection += float64(app.game.TileSize())
	}

	// closest x-coordinate to the horizontal grid intersection
	xIntersection := app.player.X() + ((yIntersection - app.player.Y()) / math.Tan(ray.Angle()))

	yStep := float64(app.game.TileSize())
	if ray.IsFacingUp() {
		yStep *= -1
	}

	xStep := float64(app.game.TileSize()) / math.Tan(ray.Angle())
	if ray.IsFacingLeft() && xStep > 0 {
		xStep *= -1
	}
	if ray.IsFacingRight() && xStep < 0 {
		xStep *= -1
	}

	nextHorzXCollision := xIntersection
	nextHorzYCollision := yIntersection

	// increments xstep and ystep until a wall collision
	for nextHorzXCollision >= 0 &&
		nextHorzXCollision <= float64(app.game.WindowWidth()) &&
		nextHorzYCollision >= 0 &&
		nextHorzYCollision <= float64(app.game.WindowHeight()) {

		xToCheck := nextHorzXCollision
		yToCheck := nextHorzYCollision
		if ray.IsFacingUp() {
			yToCheck -= 1
		}
		if app.game.IsSolidCoordinate(xToCheck, yToCheck) {
			horzXWallCollision = nextHorzXCollision
			horzYWallCollision = nextHorzYCollision
			horzWallContent = app.game.GameMap()[int(math.Floor(yToCheck/float64(app.game.TileSize())))][int(math.Floor(xToCheck/float64(app.game.TileSize())))]
			foundHorzCollision = true
			break
		} else {
			nextHorzXCollision += float64(xStep)
			nextHorzYCollision += float64(yStep)
		}
	}

	// Starting vertical calculations
	var vertXWallCollision float64
	var vertYWallCollision float64
	var vertWallContent int32
	foundVertCollision := false

	xIntersection = math.Floor(app.player.X()/float64(app.game.TileSize())) * float64(app.game.TileSize())
	if ray.IsFacingRight() {
		xIntersection += float64(app.game.TileSize())
	}

	yIntersection = app.player.Y() + ((xIntersection - app.player.X()) * math.Tan(ray.Angle()))

	xStep = float64(app.game.TileSize())
	if ray.IsFacingLeft() {
		xStep *= -1
	}

	yStep = float64(app.game.TileSize()) * math.Tan(ray.Angle())
	if ray.IsFacingUp() && yStep > 0 {
		yStep *= -1
	}
	if ray.IsFacingDown() && yStep < 0 {
		yStep *= -1
	}

	nextVertXCollision := xIntersection
	nextVertYCollision := yIntersection

	for nextVertXCollision >= 0 &&
		nextVertXCollision <= float64(app.game.WindowWidth()) &&
		nextVertYCollision >= 0 &&
		nextVertYCollision <= float64(app.game.WindowHeight()) {

		xToCheck := nextVertXCollision
		if ray.IsFacingLeft() {
			xToCheck -= 1
		}
		yToCheck := nextVertYCollision
		if app.game.IsSolidCoordinate(xToCheck, yToCheck) {
			vertXWallCollision = nextVertXCollision
			vertYWallCollision = nextVertYCollision
			vertWallContent = app.game.GameMap()[int(yToCheck/float64(app.game.TileSize()))][int(xToCheck/float64(app.game.TileSize()))]
			foundVertCollision = true
			break
		} else {
			nextVertXCollision += float64(xStep)
			nextVertYCollision += float64(yStep)
		}
	}

	horzCollisionDist := math.MaxFloat64
	vertCollisionDist := math.MaxFloat64

	if foundHorzCollision {
		horzCollisionDist = utils.DistanceBetweenPoints(app.player.X(), app.player.Y(), horzXWallCollision, horzYWallCollision)
	}
	if foundVertCollision {
		vertCollisionDist = utils.DistanceBetweenPoints(app.player.X(), app.player.Y(), vertXWallCollision, vertYWallCollision)
	}

	if vertCollisionDist < horzCollisionDist {
		app.game.Rays()[rayId].SetDistance(vertCollisionDist)
		app.game.Rays()[rayId].SetXCollision(vertXWallCollision)
		app.game.Rays()[rayId].SetYCollision(vertYWallCollision)
		app.game.Rays()[rayId].SetContent(vertWallContent)
		app.game.Rays()[rayId].SetIsVerticalCollision(true)
	} else {
		app.game.Rays()[rayId].SetDistance(horzCollisionDist)
		app.game.Rays()[rayId].SetXCollision(horzXWallCollision)
		app.game.Rays()[rayId].SetYCollision(horzYWallCollision)
		app.game.Rays()[rayId].SetContent(horzWallContent)
		app.game.Rays()[rayId].SetIsVerticalCollision(false)
	}
}

func (app *App) CastRays() {
	for col := 0; col < int(app.game.NumRays()); col++ {
		angle := app.player.RotationAngle() + math.Atan((float64(col)-float64(app.game.NumRays())/2)/float64(app.game.DistanceProjectionPlane()))
		app.CastRay(angle, col)
	}
}

func (app *App) RenderRays() {
	app.renderer.SetDrawColor(255, 0, 0, 255)
	for i := 0; i < int(app.game.NumRays()); i++ {
		app.renderer.DrawLine(
			int32(app.game.MinimapScale()*app.player.X()),
			int32(app.game.MinimapScale()*app.player.Y()),
			int32(app.game.MinimapScale()*app.game.Rays()[i].XCollision()),
			int32(app.game.MinimapScale()*app.game.Rays()[i].YCollision()),
		)
	}
}
