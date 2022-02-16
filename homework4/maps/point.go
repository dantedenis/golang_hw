package maps

type Point2d struct {
	x float64
	y float64
}

func NewPoint2d(x, y float64) *Point2d {
	return &Point2d{x: x, y: y}
}
