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
	dCoord := maps.NewArc(*pointer2, *pointer3)
	fmt.Println(dCoord)
	fmt.Printf("%T\n%T", pointer3, dCoord)
}
