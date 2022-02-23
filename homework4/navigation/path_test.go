package navigation

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"homework4/maps"
	"homework4/mock"
	"homework4/navigation/info"
	"testing"
)

func TestNavigation_PathInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGeocoder := mock.NewMockGeocoding(ctrl)
	point1 := maps.NewPointDeg(32, 45)
	point2 := maps.NewPointDeg(0, 0)
	mockGeocoder.EXPECT().ReverseGeocoder(point1).Return(info.GeocodeData{Point: *point1, City: "moscow", Postal: "123123", Country: "russia"})
	mockGeocoder.EXPECT().ReverseGeocoder(point1).Return(info.GeocodeData{Point: *point2, City: "novgorod", Postal: "141123", Country: "russia"})

	navi := NewNav(mockGeocoder)
	info, err := navi.PathInfo(*point1, *point2)
	assert.Nil(t, err)
	assert.Equal(t, info.PlaceStart().Point, *point1)
	assert.Equal(t, info.PlaceFinish().Point, *point2)
	assert.Equal(t, info.PlaceFinish().Postal, "131313")
	assert.Equal(t, info.PlaceStart().Postal, "123321")
}
