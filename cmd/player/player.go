package player

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/pedro-git-projects/go-raycasting/cmd/game"
	"github.com/veandco/go-sdl2/sdl"
)

// turnDirection is the type used in the enum for the player turn direction
type turnDirection int

// walkDirection is the type used in the enum for the player walk direction
type walkDirection int

const (
	turnNeutral turnDirection = 0
	left        turnDirection = -1
	right       turnDirection = 1
)

const (
	walkNeutral walkDirection = 0
	foward      walkDirection = 1
	backward    walkDirection = -1
)

// the player type represents the player game object
type Player struct {
	x             float64
	y             float64
	width         float64
	height        float64
	turnDirection turnDirection
	walkDirection walkDirection
	rotationAngle float64
	walkSpeed     float64
	turnSpeed     float64
	minimapScale  float64
}

// New creates a pointer to a player in the x and y postion
// with 1 pixel witdh and height
// with neutral turn and walk diretion
// 100 pixel walkspeed and 45 radians turning
// the default minimapScaling is 0.3
func New(x, y float64) *Player {
	p := &Player{
		x:             x,
		y:             y,
		width:         1,
		height:        1,
		turnDirection: turnNeutral,
		walkDirection: walkNeutral,
		rotationAngle: math.Pi / 2,
		walkSpeed:     100,
		turnSpeed:     45 * (math.Pi / 180),
		minimapScale:  game.MinimapScaling,
	}
	return p
}

/* Accesssors */

func (p Player) X() float64 {
	return p.x
}

func (p Player) Y() float64 {
	return p.y
}

func (p Player) RotationAngle() float64 {
	return p.rotationAngle
}

/* Setters */

func (p *Player) SetX(x float64) {
	p.x = x
}

func (p *Player) SetY(y float64) {
	p.y = y
}

func (p *Player) IncX(x float64) {
	p.x += x
}

func (p *Player) IncY(y float64) {
	p.y += y
}

func (p *Player) DecX(x float64) {
	p.x -= x
}

func (p *Player) DecY(y float64) {
	p.y -= y
}

// Render takes a pointer to a renderer and
// draws the player object and a guiding line
// the size is controlled by the minimapScale fieled
func (p *Player) Render(r *sdl.Renderer) {
	r.SetDrawColor(255, 255, 255, 255)
	playerRect := sdl.Rect{
		X: int32(p.x * p.minimapScale),
		Y: int32(p.y * p.minimapScale),
		W: int32(p.width * p.minimapScale),
		H: int32(p.height * p.minimapScale),
	}
	r.FillRect(&playerRect)
	length := 40
	r.DrawLine(
		int32(p.minimapScale*p.x),
		int32(p.minimapScale*p.y),
		int32(p.minimapScale*p.x+math.Cos(p.rotationAngle)*float64(length)),
		int32(p.minimapScale*p.y+math.Sin(p.rotationAngle)*float64(length)),
	)
}

// SetWalkDirection takes a string and matches it to the corresponding enum value
// the matched value is then used to set the field accordingly
// if no matc is found an error is returned
func (p *Player) SetWalkDirection(direction string) error {
	switch strings.ToLower(direction) {
	case "neutral":
		p.walkDirection = walkNeutral
	case "foward":
		p.walkDirection = foward
	case "backward":
		p.walkDirection = backward
	default:
		err := fmt.Sprintf("Unknown walk direction %s", direction)
		return errors.New(err)
	}
	return nil
}

// SetTurnirection takes a string and matches it to the corresponding enum value
// the matched value is then used to set the field accordingly
// if no matc is found an error is returned
func (p *Player) SetTurnDirection(direction string) error {
	switch strings.ToLower(direction) {
	case "neutral":
		p.turnDirection = turnNeutral
	case "left":
		p.turnDirection = left
	case "right":
		p.turnDirection = right
	default:
		err := fmt.Sprintf("Unknown walk direction %s", direction)
		return errors.New(err)
	}
	return nil
}

// Move changes the value of the x and y fiels of a player object
// to new values considering the turn and walk speeds
// accounting fot the delta time
// and checking for collision into the new position
func (p *Player) Move(delta float64, g *game.Game) {
	p.rotationAngle += float64(p.turnDirection) * p.turnSpeed * delta
	distance := float64(p.walkDirection) * p.walkSpeed * delta
	newX := p.x + math.Cos(p.rotationAngle)*distance
	newY := p.y + math.Sin(p.rotationAngle)*distance
	if !g.IsSolidCoordinate(newX, newY) {
		p.SetX(newX)
		p.SetY(newY)
	}
}
