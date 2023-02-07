package window

import "math"

const (
	MinimapScaling float64 = 0.2
	TileSize       int32   = 64
	NumRows        int32   = 13
	NumCols        int32   = 20
	FOV            float64 = 60 * (math.Pi / 180)
)

var (
	Width             int32   = NumCols * TileSize
	Height            int32   = NumRows * TileSize
	DistanceProjPlane float64 = ((float64(Width) / 2) / math.Tan(FOV/2))
	NumRays                   = Width
)
