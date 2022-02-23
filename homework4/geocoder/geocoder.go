package geocoder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"homework4/maps"
	"homework4/navigation/info"
	"net/http"
)

type Geocoding interface {
	ReverseGeocoder(planet maps.PointPlanet) (info.GeocodeData, error)
}

type Geocoder struct {
	client *http.Client
	url    string
	token  string
	user   string
	secret string
}

func NewGeocoder(url, token, user, secret string) *Geocoder {
	return &Geocoder{url: url, token: token, client: &http.Client{}, user: user, secret: secret}
}

func (g Geocoder) ReverseGeocoder(point maps.PointPlanet) (data info.GeocodeData, err error) {
	jsonRequest, err := json.Marshal(map[string]string{"lat": fmt.Sprintf("%f", point.LatDeg()), "lon": fmt.Sprintf("%f", point.LngDeg()), "radius_meters": "10"})
	if err != nil {
		return data, errors.New("error jsonMarshall")
	}
	req, err := http.NewRequest("POST", g.url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return data, errors.New("init request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Token "+g.token)
	response, err := g.client.Do(req)
	if err != nil {
		return data, errors.New("send POST on the server")
	}
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return data, errors.New("decode json")
	}
	suggestions := result["suggestions"].([]interface{})[0].(map[string]interface{})
	dataMap := suggestions["data"].(map[string]interface{})
	pointStr, err := maps.NewPoint(dataMap["geo_lat"].(string), dataMap["geo_lon"].(string))
	if err != nil {
		return data, err
	}
	data.City = dataMap["region"].(string)
	data.Point = *pointStr
	data.Country = dataMap["country"].(string)
	return
}

func (g Geocoder) Geocoding(str string) (data info.GeocodeData, err error) {
	jsonRequest, _ := json.Marshal([]string{str})
	req, err := http.NewRequest("POST", g.url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return data, errors.New("init request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+g.token)
	req.Header.Add("X-Secret", g.secret)
	response, err := g.client.Do(req)
	if err != nil {
		return data, errors.New("send POST on the server")
	}
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return data, errors.New("decode json")
	}
	fmt.Println(result)
	return
}
