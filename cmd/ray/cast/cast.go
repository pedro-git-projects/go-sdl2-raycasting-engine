package cast

import (
	"math"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/pedro-git-projects/go-raycasting/cmd/player"
	"github.com/pedro-git-projects/go-raycasting/cmd/ray"
	"github.com/veandco/go-sdl2/sdl"
)

func CastRay(angle float64, rayId int, g *game.Game, p *player.Player) {
	g.Rays()[rayId] = *ray.New(angle)
	ray := g.Rays()[rayId]
	// closest y-coordinate to the horizontal grid intersection
	yIntersection := math.Floor(p.Y()/float64(g.TileSize())) * float64(g.TileSize())
	if ray.IsFacingDown() {
		yIntersection += float64(g.TileSize())
	}

	// closest x-coordinate to the horizontal grid intersection
	xIntersection := p.X() + (yIntersection-p.Y())/math.Tan(ray.Angle())

	var horzXWallCollision float64
	var horzYWallCollision float64
	var horzWallContent int32
	foundHorzCollision := false

	yStep := g.TileSize()
	if ray.IsFacingUp() {
		yStep *= -1
	}

	xStep := g.TileSize() / int32(math.Tan(ray.Angle()))
	if ray.IsFacingLeft() && xStep > 0 {
		xStep *= -1
	}
	if ray.IsFacingRight() && xStep < 0 {
		xStep *= -1
	}

	nextHorzXCollision := xIntersection
	nextHorzYCollision := yIntersection

	// increments xstep and ystep until a wall collision
	for nextHorzXCollision >= 0 && nextHorzXCollision <= float64(g.Cols())*float64(g.TileSize()) && nextHorzYCollision >= 0 && nextHorzYCollision <= float64(g.Rows())*float64(g.TileSize()) {
		xToCheck := nextHorzXCollision
		yToCheck := nextHorzYCollision
		if ray.IsFacingUp() {
			yToCheck -= 1
		}
		if g.IsSolidCoordinate(xToCheck, yToCheck) {
			horzXWallCollision = nextHorzXCollision
			horzYWallCollision = nextHorzYCollision
			horzWallContent = g.GameMap()[int32(math.Floor(yToCheck/float64(g.TileSize())))][int32(math.Floor(xToCheck/float64(g.TileSize())))]
			foundHorzCollision = true
			break
		} else {
			nextHorzXCollision += float64(xStep)
			nextHorzYCollision += float64(yStep)
		}
	}

	xIntersection = math.Floor(p.X()/float64(g.TileSize())) * float64(g.TileSize())
	if ray.IsFacingRight() {
		xIntersection += float64(g.TileSize())
	}

	yIntersection = p.X() + (xIntersection-p.X())*math.Tan(ray.Angle())

	var vertXWallCollision float64
	var vertYWallCollision float64
	var vertWallContent int32
	foundVertCollision := false

	xStep = g.TileSize()
	if ray.IsFacingLeft() {
		xStep *= -1
	}

	yStep = g.TileSize() * int32(math.Tan(ray.Angle()))
	if ray.IsFacingUp() && yStep > 0 {
		yStep *= -1
	}
	if ray.IsFacingDown() && yStep < 0 {
		yStep *= -1
	}

	nextVertXCollision := xIntersection
	nextVertYCollision := yIntersection

	for nextVertXCollision >= 0 && nextVertXCollision <= float64(g.Cols())*float64(g.TileSize()) && nextVertYCollision >= 0 && nextVertYCollision <= float64(g.Rows())*float64(g.TileSize()) {
		xToCheck := nextVertXCollision
		if ray.IsFacingLeft() {
			xToCheck -= 1
		}
		yToCheck := nextVertYCollision
		if g.IsSolidCoordinate(xToCheck, yToCheck) {
			vertXWallCollision = nextVertXCollision
			vertYWallCollision = nextVertYCollision
			vertWallContent = g.GameMap()[int32(yToCheck/float64(g.TileSize()))][int32(xToCheck/float64(g.TileSize()))]
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
		horzCollisionDist = distanceBetweenPoints(p.X(), p.Y(), horzXWallCollision, horzYWallCollision)
	}
	if foundVertCollision {
		vertCollisionDist = distanceBetweenPoints(p.X(), p.Y(), vertXWallCollision, vertYWallCollision)
	}

	if vertCollisionDist < horzCollisionDist {
		g.Rays()[rayId].SetDistance(vertCollisionDist)
		g.Rays()[rayId].SetXCollision(vertXWallCollision)
		g.Rays()[rayId].SetYCollision(vertYWallCollision)
		g.Rays()[rayId].SetContent(vertWallContent)
		g.Rays()[rayId].SetIsVerticalCollision(true)
	} else {
		g.Rays()[rayId].SetDistance(horzCollisionDist)
		g.Rays()[rayId].SetXCollision(horzXWallCollision)
		g.Rays()[rayId].SetYCollision(horzYWallCollision)
		g.Rays()[rayId].SetContent(horzWallContent)
		g.Rays()[rayId].SetIsVerticalCollision(false)
	}
}

func distanceBetweenPoints(x0, x1, y0, y1 float64) float64 {
	return math.Sqrt((x1-x0)*(x1-x0) + (y1-y0)*(y1-y0))
}

func CastRays(g *game.Game, p *player.Player) {
	for col := 0; col < int(g.NumRays()); col++ {
		// TODO: player rotation angle is being divided by zero
		// find where
		angle := p.RotationAngle()
		// p.RotationAngle() +	math.Atan((float64(col) - float64(g.NumRays())/2) / float64(g.DistanceProjectionPlane())))
		CastRay(angle, col, g, p)
	}
}

func RenderRays(r *sdl.Renderer, g *game.Game, p *player.Player) {
	r.SetDrawColor(255, 0, 0, 255)
	for i := 0; i < int(g.NumRays()); i++ {
		r.DrawLine(
			int32(g.MinimapScale()*p.X()),
			int32(g.MinimapScale()*p.Y()),
			int32(g.MinimapScale()*g.Rays()[i].XCollision()),
			int32(g.MinimapScale()*g.Rays()[i].YCollision()),
		)
	}
}
