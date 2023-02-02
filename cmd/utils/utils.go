package utils

import "math"

func DistanceBetweenPoints(x0, x1, y0, y1 float64) float64 {
	return math.Sqrt((x1-x0)*(x1-x0) + (y1-y0)*(y1-y0))
}
