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

/*
curl -X POST \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -H "Authorization: Token d263a7aad9376d367f7efa7b55133f90a006a71e" \
  -d '{ "lat": 55.878, "lon": 37.653 }' \
  https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address
*/

func (g Geocoder) ReverseGeocoder(point maps.PointPlanet) (data info.GeocodeData, err error) {
	jsonRequest, _ := json.Marshal(map[string]string{"lat": fmt.Sprintf("%f", 55.878), "lon": fmt.Sprintf("%f", 37.653), "radius_meters": "50"})
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
	//buf, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(buf))
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return data, errors.New("decode json")
	}
	fmt.Println(result)
	return
}

/*
curl -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Token d263a7aad9376d367f7efa7b55133f90a006a71e" \
    -H "X-Secret: 75f146fc3c58b5ca254382029a7a440a882c5e79" \
    -d '[ "москва сухонская 11" ]' \
    https://cleaner.dadata.ru/api/v1/clean/address
*/
func (g Geocoder) Geocoding(str string) (data info.GeocodeData, err error) {
	jsonRequest, _ := json.Marshal(str)
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
	fmt.Println(response)
	var result map[string]string
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return data, errors.New("decode json")
	}
	fmt.Println(result)
	return
}

func (g Geocoder) RevGeo(point maps.PointPlanet) (data info.GeocodeData, err error) {
	jsonRequest, _ := json.Marshal(map[string]string{"latitude": "43", "longitude": "30"})
	req, err := http.NewRequest("POST", g.url, bytes.NewBufferString(string(jsonRequest)))
	if err != nil {
		return data, errors.New("init request")
	}
	req.Header.Add("user-id", g.user)
	req.Header.Add("api-key", g.token)
	//req.Header.Add("url-info", g.url)
	//req.Header.Add("output-format", "JSON")
	//req.Header.Add("output-case", "kebab")
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

/*
func (g Geocoder) GeoAddress(str string) (data info.GeocodeData, err error) {
	jsonRequest, _ := json.Marshal(map[string]string{"city": "moscow"})
	req, err := http.NewRequest("POST", g.url, bytes.NewBufferString())
}
*/
