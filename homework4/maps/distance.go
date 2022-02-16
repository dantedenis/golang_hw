package maps

type Coordinate interface {
	Distance() float64
}

type Arc struct {
	start PointPlanet
	end   PointPlanet
	long  float64
	name  string
}

func NewArc(p1, p2 PointPlanet) *Arc {
	return &Arc{start: p1, end: p2, long: p1.Distance(p2)}
}

func (a Arc) Distance() float64 {
	return a.start.Distance(a.end)
}
