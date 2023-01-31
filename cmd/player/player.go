package player

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// TODO fix palyer rendering

type turnDirection int
type walkDirection int

const (
	turnNeutral turnDirection = 0
	left        turnDirection = 1
	right       turnDirection = -1
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
		minimapScale:  0.3,
	}
	return p
}

func (p Player) X() float64 {
	return p.x
}

func (p Player) Y() float64 {
	return p.y
}

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

func (p *Player) Render(r *sdl.Renderer) {
	r.SetDrawColor(255, 255, 255, 255)
	playerRect := sdl.Rect{
		X: int32(p.x * p.minimapScale),
		Y: int32(p.y * p.minimapScale),
		W: int32(p.width * p.minimapScale),
		H: int32(p.height * p.minimapScale),
	}
	r.FillRect(&playerRect)
	r.DrawLine(
		int32(p.minimapScale*p.x),
		int32(p.minimapScale*p.y),
		int32(p.minimapScale*(p.x+math.Cos(p.rotationAngle)*40)),
		int32(p.minimapScale*(p.x+math.Sin(p.rotationAngle)*40)),
	)
}

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

func (p *Player) Move() {
	p.rotationAngle = float64(p.turnDirection) * p.turnSpeed
	distance := float64(p.walkDirection) * p.walkSpeed
	newX := p.x + math.Cos(p.rotationAngle)*distance
	newY := p.y + math.Sin(p.rotationAngle)*distance
	p.SetX(newX)
	p.SetY(newY)
}
