package navigation

type AllPath interface {
	Distance() float64
}

type Navigation struct {
	distance []AllPath
}

func NewNav(dist ...AllPath) *Navigation {
	return &Navigation{distance: dist}
}

func (nav Navigation) GetDist() (dist float64) {
	for _, d := range nav.distance {
		temp := d.Distance()
		dist += temp
	}
	return
}
