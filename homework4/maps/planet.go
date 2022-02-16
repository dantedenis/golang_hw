package maps

import "math"

const (
	oneRadian   Degree  = math.Pi / 180
	radiusEarth float64 = 6372795
)

type Degree float64
type Radian float64

type PointPlanet struct {
	lat Radian
	lng Radian
}

func NewPointDeg(latitude, longitude Degree) *PointPlanet {
	return &PointPlanet{lat: toRadian(latitude), lng: toRadian(longitude)}
}

func NewPointRad(latitude, longitude Radian) *PointPlanet {
	return &PointPlanet{lat: latitude, lng: longitude}
}

func toRadian(value Degree) Radian {
	return Radian(value * oneRadian)
}

func (p PointPlanet) Lat() Radian {
	return p.lat
}

func (p PointPlanet) Lng() Radian {
	return p.lng
}

func deltaLng(p1 PointPlanet, p2 PointPlanet) Radian {
	return p2.lng - p1.lng
}

func cos(value Radian) float64 {
	return math.Cos(float64(value))
}

func sin(value Radian) float64 {
	return math.Sin(float64(value))
}

func (p1 PointPlanet) Distance(p2 PointPlanet) float64 {
	Lat1 := p1.Lat()
	Lat2 := p2.Lat()
	deltaLn := deltaLng(p1, p2)

	y := math.Sqrt(math.Pow(cos(Lat2)*sin(deltaLn), 2) +
		math.Pow(cos(Lat1)*sin(Lat2)-sin(Lat1)*cos(Lat2)*cos(deltaLn), 2))
	x := sin(Lat1)*sin(Lat2) + cos(Lat1)*cos(Lat2)*cos(deltaLn)
	return math.Atan2(y, x) * radiusEarth
}
