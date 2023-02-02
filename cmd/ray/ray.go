package ray

import (
	"math"

	"github.com/pedro-git-projects/go-raycasting/cmd/utils"
)

// Ray represents a casted ray
type Ray struct {
	angle               float64
	xCollision          float64
	yCollision          float64
	distance            float64
	isVerticalCollision bool
	isFacingUp          bool
	isFacingDown        bool
	isFacingLeft        bool
	isFacingRight       bool
	content             int32
}

// New creates a a pointer to a ray with normalized angle and booleans to its facing diretction
func New(angle float64) *Ray {
	utils.NormalizeAngle(&angle)
	down := isFacingDown(angle)
	up := !down
	right := isFacingRight(angle)
	left := !right
	r := &Ray{
		angle:         angle,
		isFacingUp:    up,
		isFacingDown:  down,
		isFacingRight: right,
		isFacingLeft:  left,
	}
	return r
}

/* Accessors  */

func (r Ray) IsFacingDown() bool {
	return r.isFacingDown
}

func (r Ray) IsFacingUp() bool {
	return r.isFacingUp
}

func (r Ray) IsFacingLeft() bool {
	return r.isFacingLeft
}

func (r Ray) IsFacingRight() bool {
	return r.isFacingRight
}

func (r Ray) Angle() float64 {
	return r.angle
}

func (r Ray) XCollision() float64 {
	return r.xCollision
}

func (r Ray) YCollision() float64 {
	return r.yCollision
}

func (r Ray) IsVerticalCollision() bool {
	return r.isVerticalCollision
}

func (r Ray) Distance() float64 {
	return r.distance
}

/* Mutators */

func (r *Ray) SetAngle(angle float64) {
	r.angle = angle
}

func (r *Ray) SetXCollision(xCollision float64) {
	r.xCollision = xCollision
}

func (r *Ray) SetYCollision(yCollision float64) {
	r.yCollision = yCollision
}

func (r *Ray) SetDistance(distance float64) {
	r.distance = distance
}

func (r *Ray) SetIsVerticalCollision(isVerticalCollision bool) {
	r.isVerticalCollision = isVerticalCollision
}

func (r *Ray) SetContent(content int32) {
	r.content = content
}

// isFacingDown returns true if an angle is facing up, false otherwise
func isFacingDown(angle float64) bool {
	if angle > 0 && angle < math.Pi {
		return true
	}
	return false
}

// isFacingRight returns true if an angle is facing right, false otherwise
func isFacingRight(angle float64) bool {
	if angle < 0.5*math.Pi || angle > 1.5*math.Pi {
		return true
	}
	return false
}
