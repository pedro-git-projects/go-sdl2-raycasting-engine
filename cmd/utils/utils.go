package utils

import "math"

// DistanceBetweenPoints calculates the distance between two 2D coordiantes
func DistanceBetweenPoints(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt((x1-x0)*(x1-x0) + (y1-y0)*(y1-y0))
}

// NormalizeAngle takes a pointer to a possibliy negative
// angle in radians and mutates it to the absolute angle
func NormalizeAngle(angle *float64) {
	*angle = math.Remainder(*angle, 2*math.Pi)
	if *angle < 0 {
		*angle = (2 * math.Pi) + *angle
	}
}
