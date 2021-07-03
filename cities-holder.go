package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"io/ioutil"
)

var Cities CitiesResponse

func startUpCities() {
	data, err := ioutil.ReadFile("./vkcitycodes.json")
	check(err)
	json.Unmarshal(data, &Cities)
}

func getCityName(id int) string {
	resp, err := resty.R().Get(urlCity(id))
	check(err)
	var city CityResponse
	json.Unmarshal([]byte(resp.String()), &city)
	if city.Body == nil {
		return ""
	}
	return city.Body[0].Name
}

func urlCity(cityid int) string {
	return fmt.Sprintf("https://api.vk.com/method/database.getCitiesById?city_ids=%v&access_token=%v&v=V", cityid, getToken())
}
