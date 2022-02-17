package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Sides int

const (
	SidesSquare   Sides = 4
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	var area float64
	switch sidesNum {
	case SidesCircle:
		area = (math.Pow(sideLen, 2)) * math.Pi
	case SidesTriangle:
		area = (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	case SidesSquare:
		area = math.Pow(sideLen, 2)
	default:
		area = 0
	}

	return area
}
