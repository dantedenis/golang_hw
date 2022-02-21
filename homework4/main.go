package main

import (
	"fmt"
	"homework4/maps"
)

func main() {
	//pointer1 := maps.NewPoint2d(0, 2)
	pointer2 := maps.NewPointDeg(2.5, 3.1231)
	pointer3 := maps.NewPointDeg(45.12, 67.543)
	//fmt.Println(pointer1.Distance(struct{ x, y float64 }{3, 5}))
	//fmt.Printf("%10f\n", pointer2.Distance(*pointer3))
	//fmt.Printf("%10f\n", pointer2.Distance(struct{ lat, lng float64 }{45.12, 67.543}))
	dCoords := maps.NewArc(*pointer2, *pointer3)
	Path := dCoords.AddPoint(*maps.NewPointDeg(13.121, 466.112), *maps.NewPointRad(1231.11, 63.21))
	fmt.Println(dCoords)
	fmt.Printf("%T\n%T", pointer3, dCoords)
	fmt.Println(Path)
}
