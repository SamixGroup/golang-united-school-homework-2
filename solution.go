package square

import "math"

// CustomInt Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type CustomInt int

const (
	SidesTriangle CustomInt = 3
	SidesSquare   CustomInt = 4
	SidesCircle   CustomInt = 0
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	switch sidesNum {
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		var sum = (3 * sideLen) / 2
		return math.Sqrt(sum * math.Pow((sum - sideLen),3))
	default:
		return 0
	}

}
