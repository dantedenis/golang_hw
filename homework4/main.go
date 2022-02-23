package main

import (
	"fmt"
	"homework4/geocoder"
	"homework4/maps"
)

func main() {
	pointer2 := maps.NewPointDeg(55.601983, 37.359486)
	pointer3 := maps.NewPointDeg(45.12, 67.543)
	dCoords := maps.NewArc(*pointer2, *pointer3)
	Path := dCoords.AddPoint(*maps.NewPointRad(13.121, 466.112), *maps.NewPointRad(1231.11, 63.21))

	fmt.Println(*dCoords)
	fmt.Printf("%T\n%T\n", pointer3, dCoords)
	fmt.Println("Before added:")
	for _, p := range Path.GetArcs() {
		fmt.Println(p)
	}
	fmt.Println("After added:")
	Path.AddPointSlice(*maps.NewPointRad(17.121, 46.112))
	for _, p := range Path.GetArcs() {
		fmt.Println(p)
	}
	fmt.Println("----------------------------------------------")
	geo := geocoder.NewGeocoder("https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", "d263a7aad9376d367f7efa7b55133f90a006a71e", "", "")
	data, err := geo.ReverseGeocoder(*pointer2)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
	fmt.Println(data)
}
