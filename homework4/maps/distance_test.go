package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	point1 PointPlanet
	point2 PointPlanet
}

type testCase struct {
	name string
	par  param
	want float64
}

func TestArc_Distance(t *testing.T) {
	t.Run("Distance", func(t *testing.T) {
		point1, point2 := PointPlanet{toRadian(12), toRadian(13)}, PointPlanet{lat: toRadian(45), lng: toRadian(45)}
		got := point1.Distance(point2)
		want := 4760627.0
		res := got - want
		if res/want >= 0.001 {
			t.Errorf("got %g want %g", got, want)
		}
	})

	tests := []testCase{
		{"1st TEST:", param{point1: PointPlanet{0.22, 0.2}, point2: PointPlanet{0.78, 0.78}}, 4760627.0},
		{"2nd TEST", param{point1: PointPlanet{0.26, 0.5}, point2: PointPlanet{0.26, 0.5}}, 0},
		{"3rd TEST", param{point1: PointPlanet{0.44, 0.01}, point2: PointPlanet{0.56, 0.78}}, 0},
	}
	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			line := NewArc(ts.par.point1, ts.par.point2)
			dist := line.Distance()
			assert.Equal(t, dist, ts.want)
		})
	}
}
