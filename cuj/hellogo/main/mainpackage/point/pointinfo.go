package point

import "math"

type Point struct {
	X float64
	Y float64
}

func Discount(p , q Point) float64{
	return math.Hypot(p.X-q.X,p.Y-q.Y)
}
