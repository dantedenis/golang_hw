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

func (a ArcSlice) GetArcs() []Arc {
	return a.arcs
}

func (a Arc) AddPoint(points ...PointPlanet) *ArcSlice {
	var result ArcSlice
	result.arcs = append(result.arcs, a)
	for _, point := range points {
		indexLast := len(result.arcs) - 1
		result.arcs = append(result.arcs, *NewArc(result.arcs[indexLast].end, point))
	}
	return &result
}

func (a *ArcSlice) AddPointSlice(points ...PointPlanet) {
	for _, point := range points {
		indexLast := len(a.arcs) - 1
		a.arcs = append(a.arcs, *NewArc(a.arcs[indexLast].end, point))
	}
}
