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

func (a Arc) AddPoint(points ...PointPlanet) result *ArcSlice {
	result = append(result, Arc)
	for _, point range points {
		indexLast := len(result) - 1
		result = append(result, NewArc(result[indexLast], point))
	}
	return &result
}

func (aSlice *ArcSlice) AddPointSlice(points ...PointPlanet) *ArcSlice {
	for _, point range points {
		indexLast := len(aSlice.arcs) - 1
		aSlice.arcs = append(aSlice.arcs, NewArc(aSlice.arcs[indexLast] ,point))
	}
	return aSlice
}
