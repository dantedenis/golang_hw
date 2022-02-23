package navigation

import (
	"errors"
	"homework4/geocoder"
	"homework4/maps"
	"homework4/navigation/info"
)

type AllPath interface {
	Distance() float64
}

type Navigation struct {
	geocoder geocoder.Geocoder
}

type PathInfo struct {
	pStart  info.GeocodeData
	pFinish info.GeocodeData
}

func (p PathInfo) PlaceStart() info.GeocodeData {
	return p.pStart
}

func (p PathInfo) PlaceFinish() info.GeocodeData {
	return p.pFinish
}

func NewNav(geocoder geocoder.Geocoder) *Navigation {
	return &Navigation{geocoder: geocoder}
}
func (n Navigation) PathInfo(p1, p2 maps.PointPlanet) (info PathInfo, err error) {
	data1, err := n.geocoder.ReverseGeocoder(p1)
	if err != nil {
		return info, errors.New("geocoder point1")
	}
	data2, err := n.geocoder.ReverseGeocoder(p2)
	if err != nil {
		return info, errors.New("geocoder point2")
	}
	info = PathInfo{pStart: data1, pFinish: data2}
	return
}
