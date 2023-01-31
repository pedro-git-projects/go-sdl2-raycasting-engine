package player

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type turnDirection int
type walkDirection int

const (
	turnNeutral turnDirection = iota
	left
	right
)

const (
	walkNeutral walkDirection = iota
	foward
	backward
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
		width:         5,
		height:        5,
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
}
