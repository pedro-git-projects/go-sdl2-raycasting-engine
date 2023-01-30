package player

// the player type represents the player game object
type Player struct {
	x int
	y int
}

func New(x, y int) *Player {
	p := &Player{
		x: x,
		y: y,
	}
	return p
}

func (p Player) X() int {
	return p.x
}

func (p Player) Y() int {
	return p.y
}

func (p *Player) SetX(x int) {
	p.x = x
}

func (p *Player) SetY(y int) {
	p.y = y
}

func (p *Player) IncX(x int) {
	p.x += x
}

func (p *Player) IncY(y int) {
	p.y += y
}

func (p *Player) DecX(x int) {
	p.x -= x
}

func (p *Player) DecY(y int) {
	p.y -= y
}
