package main

import (
	"fmt"
	"homework4/geocoder"
	"homework4/maps"
)

func main() {
	pointer2 := maps.NewPointRad(55.878, 37.653)
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
	_, err := geo.ReverseGeocoder(*pointer2)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
	fmt.Println("----------------------------------------------")
	geo1 := geocoder.NewGeocoder("https://neutrinoapi.net/geocode-reverse", "RXlpjPVS160ql77KqzMkuRSib244EHIZXm6iJza7Ji9umj4Z", "dAnte", "")
	_, err = geo1.RevGeo(*pointer3)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
	fmt.Println("----------------------------------------------")
	geo2 := geocoder.NewGeocoder("https://cleaner.dadata.ru/api/v1/clean/address", "d263a7aad9376d367f7efa7b55133f90a006a71e", "", "75f146fc3c58b5ca254382029a7a440a882c5e79")
	_, err = geo2.Geocoding("москва сухонская 11")
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
}
