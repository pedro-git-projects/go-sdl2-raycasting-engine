package main

import (
	"math"

	"github.com/pedro-git-projects/go-raycasting/cmd/ray"
	"github.com/pedro-git-projects/go-raycasting/cmd/utils"
	"github.com/pedro-git-projects/go-raycasting/cmd/window"
)

// horizontalIntersectionResult represents the data expected when calculating the horizontal intersection
type horizontalIntersectionResult struct {
	horzXWallCollision float64
	horzYWallCollision float64
	horzWallContent    int32
	foundHorzCollision bool
}

// verticalIntersectionResult represents the data expected when calculating the vertical intersection
type verticalIntersectionResult struct {
	vertXWallCollision float64
	vertYWallCollision float64
	vertWallContent    int32
	foundVertCollision bool
}

func (app *App) calculateHorizontalIntersection(ray *ray.Ray) *horizontalIntersectionResult {
	r := &horizontalIntersectionResult{}

	// closest y-coordinate to the horizontal grid intersection
	yIntersection := math.Floor(app.player.Y()/float64(app.game.TileSize())) * float64(app.game.TileSize()) //ok
	if ray.IsFacingDown() {
		yIntersection += float64(app.game.TileSize()) // ok
	}

	// closest x-coordinate to the horizontal grid intersection
	xIntersection := app.player.X() + ((yIntersection - app.player.Y()) / math.Tan(ray.Angle())) // ok

	yStep := float64(app.game.TileSize()) // ok
	if ray.IsFacingUp() {
		yStep *= -1
	}

	xStep := float64(app.game.TileSize()) / math.Tan(ray.Angle()) // ok
	if ray.IsFacingLeft() && xStep > 0 {
		xStep *= -1
	}
	if ray.IsFacingRight() && xStep < 0 {
		xStep *= -1
	}

	// ok
	nextHorzXCollision := xIntersection
	nextHorzYCollision := yIntersection

	// increments xstep and ystep until a wall collision
	for nextHorzXCollision >= 0 &&
		nextHorzXCollision <= float64(window.Width) &&
		nextHorzYCollision >= 0 &&
		nextHorzYCollision <= float64(window.Height) {

		xToCheck := nextHorzXCollision
		yToCheck := nextHorzYCollision
		if ray.IsFacingUp() {
			yToCheck -= 1
		}
		if app.game.IsSolidCoordinate(xToCheck, yToCheck) {
			r.horzXWallCollision = nextHorzXCollision
			r.horzYWallCollision = nextHorzYCollision
			r.horzWallContent = app.game.GameMap()[int(math.Floor(yToCheck/float64(app.game.TileSize())))][int(math.Floor(xToCheck/float64(app.game.TileSize())))]
			r.foundHorzCollision = true
			break
		} else {
			nextHorzXCollision += float64(xStep)
			nextHorzYCollision += float64(yStep)
		}
	}
	return r
}

func (app *App) calculateVerticalIntersection(ray *ray.Ray) *verticalIntersectionResult {
	r := &verticalIntersectionResult{}
	xIntersection := math.Floor(app.player.X()/float64(app.game.TileSize())) * float64(app.game.TileSize())
	if ray.IsFacingRight() {
		xIntersection += float64(app.game.TileSize())
	}

	yIntersection := app.player.Y() + ((xIntersection - app.player.X()) * math.Tan(ray.Angle()))

	xStep := float64(app.game.TileSize())
	if ray.IsFacingLeft() {
		xStep *= -1
	}

	yStep := float64(app.game.TileSize()) * math.Tan(ray.Angle())
	if ray.IsFacingUp() && yStep > 0 {
		yStep *= -1
	}
	if ray.IsFacingDown() && yStep < 0 {
		yStep *= -1
	}

	nextVertXCollision := xIntersection
	nextVertYCollision := yIntersection

	for nextVertXCollision >= 0 &&
		nextVertXCollision <= float64(window.Width) &&
		nextVertYCollision >= 0 &&
		nextVertYCollision <= float64(window.Height) {

		xToCheck := nextVertXCollision
		if ray.IsFacingLeft() {
			xToCheck -= 1
		}
		yToCheck := nextVertYCollision
		if app.game.IsSolidCoordinate(xToCheck, yToCheck) {
			r.vertXWallCollision = nextVertXCollision
			r.vertYWallCollision = nextVertYCollision
			r.vertWallContent = app.game.GameMap()[int(yToCheck/float64(app.game.TileSize()))][int(xToCheck/float64(app.game.TileSize()))]
			r.foundVertCollision = true
			break
		} else {
			nextVertXCollision += float64(xStep)
			nextVertYCollision += float64(yStep)
		}
	}
	return r
}

func (app *App) CastRay(angle float64, rayId int) {
	app.game.Rays()[rayId] = *ray.New(angle)
	ray := app.game.Rays()[rayId]

	h := app.calculateHorizontalIntersection(&ray)
	v := app.calculateVerticalIntersection(&ray)

	horzCollisionDist := math.MaxFloat64
	vertCollisionDist := math.MaxFloat64

	if h.foundHorzCollision {
		horzCollisionDist = utils.DistanceBetweenPoints(app.player.X(), app.player.Y(), h.horzXWallCollision, h.horzYWallCollision)
	}
	if v.foundVertCollision {
		vertCollisionDist = utils.DistanceBetweenPoints(app.player.X(), app.player.Y(), v.vertXWallCollision, v.vertYWallCollision)
	}

	if vertCollisionDist < horzCollisionDist {
		app.game.Rays()[rayId].SetDistance(vertCollisionDist)
		app.game.Rays()[rayId].SetXCollision(v.vertXWallCollision)
		app.game.Rays()[rayId].SetYCollision(v.vertYWallCollision)
		app.game.Rays()[rayId].SetContent(v.vertWallContent)
		app.game.Rays()[rayId].SetIsVerticalCollision(true)
	} else {
		app.game.Rays()[rayId].SetDistance(horzCollisionDist)
		app.game.Rays()[rayId].SetXCollision(h.horzXWallCollision)
		app.game.Rays()[rayId].SetYCollision(h.horzYWallCollision)
		app.game.Rays()[rayId].SetContent(h.horzWallContent)
		app.game.Rays()[rayId].SetIsVerticalCollision(false)
	}
}

func (app *App) CastRays() {
	for col := 0; col < int(window.NumRays); col++ {
		angle := app.player.RotationAngle() + math.Atan((float64(col)-float64(window.NumRays)/2)/float64(window.DistanceProjPlane))
		app.CastRay(angle, col)
	}
}

func (app *App) RenderRays() {
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
