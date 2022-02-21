package maps

type Arc struct {
	start PointPlanet
	end   PointPlanet
	long  float64
	name  string
}

type ArcSlice struct {
	arcs []Arc
}

func NewArc(p1, p2 PointPlanet) *Arc {
	return &Arc{start: p1, end: p2, long: p1.Distance(p2)}
}

func (a Arc) Distance() float64 {
	return a.long
}

func (a Arc) AddPoint(points ...PointPlanet) (result *ArcSlice) {
	result.arcs = append(result.arcs, a)
	for _, point := range points {
		indexLast := len(result.arcs) - 1
		result.arcs = append(result.arcs, *NewArc(result.arcs[indexLast].end, point))
	}
	return
}

func (aSlice *ArcSlice) AddPointSlice(points ...PointPlanet) *ArcSlice {
	for _, point := range points {
		indexLast := len(aSlice.arcs) - 1
		aSlice.arcs = append(aSlice.arcs, *NewArc(aSlice.arcs[indexLast].end, point))
	}
	return aSlice
}
