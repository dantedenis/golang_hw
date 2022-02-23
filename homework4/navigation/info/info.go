package info

import "homework4/maps"

type Geocoding interface {
	Geocode(address string) (point maps.PointPlanet, err error)
	ReverseGeocode(point maps.PointPlanet) (data GeocodeData, err error)
}

type GeocodeData struct {
	Point   maps.PointPlanet
	City    string
	Postal  string `json:"postal"`
	Country string `json:"country"`
}
