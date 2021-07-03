package main

import (
	"encoding/json"
	"encoding/xml"

	//"encoding/xml"
	"fmt"
	"gopkg.in/resty.v1"
	"strconv"
	"strings"
)

const VERSION = "5.131"

func friendsInfo(userid int, count int) FriendsResponse {
	resp, err := resty.R().Get(urlFriends(userid, count))
	//fmt.Println(resp)
	check(err)
	checkResp(resp)
	var friends FriendsResponse
	//resp := "{\"response\":[{\"uid\":165062,\"first_name\":\"Vitaly\",\"last_name\":\"Skripkin\",\"bdate\":\"6.11\",\"city\":1,\"user_id\":165062},{\"uid\":230690,\"first_name\":\"Ochir\",\"last_name\":\"Lazarev\",\"bdate\":\"31.12.1989\",\"city\":1,\"user_id\":230690},{\"uid\":579802,\"first_name\":\"Artyom\",\"last_name\":\"Grachyov\",\"bdate\":\"27.2\",\"city\":1,\"user_id\":579802},{\"uid\":629638,\"first_name\":\"Tolya\",\"last_name\":\"Dzhudzhiev\",\"city\":166,\"user_id\":629638},{\"uid\":663705,\"first_name\":\"Vladimir\",\"last_name\":\"Lidzhiev\",\"bdate\":\"16.7\",\"city\":0,\"user_id\":663705},{\"uid\":745020,\"first_name\":\"Adyan\",\"last_name\":\"Erktn\",\"city\":166,\"user_id\":745020},{\"uid\":789174,\"first_name\":\"Maria\",\"last_name\":\"Romanova\",\"city\":0,\"user_id\":789174},{\"uid\":1046758,\"first_name\":\"Tserya\",\"last_name\":\"Kyurgn\",\"bdate\":\"24.2\",\"city\":1,\"user_id\":1046758},{\"uid\":1081855,\"first_name\":\"Elena\",\"last_name\":\"Fontalina\",\"bdate\":\"21.4.1990\",\"city\":1,\"user_id\":1081855},{\"uid\":1666211,\"first_name\":\"Aldar\",\"last_name\":\"Manzhikov\",\"bdate\":\"9.4\",\"city\":166,\"user_id\":1666211},{\"uid\":1894027,\"first_name\":\"Erdnya\",\"last_name\":\"Kinzeev\",\"bdate\":\"4.12.1993\",\"city\":1,\"user_id\":1894027},{\"uid\":2006446,\"first_name\":\"Maria\",\"last_name\":\"Tikhanovskaya\",\"bdate\":\"13.11.1994\",\"city\":0,\"user_id\":2006446},{\"uid\":2114581,\"first_name\":\"Evgenia\",\"last_name\":\"Makedonskaya\",\"bdate\":\"10.12.1993\",\"city\":1,\"user_id\":2114581},{\"uid\":2204562,\"first_name\":\"Artyom\",\"last_name\":\"Vinokurov\",\"bdate\":\"5.3\",\"city\":1,\"user_id\":2204562},{\"uid\":2310589,\"first_name\":\"Ochir\",\"last_name\":\"Shuraev\",\"bdate\":\"2.6\",\"city\":1,\"user_id\":2310589},{\"uid\":2516420,\"first_name\":\"Andrey\",\"last_name\":\"Maximov\",\"bdate\":\"14.4.1991\",\"city\":1,\"user_id\":2516420},{\"uid\":2556392,\"first_name\":\"Sergey\",\"last_name\":\"Nemeshaev\",\"bdate\":\"29.1.1990\",\"city\":1,\"user_id\":2556392},{\"uid\":2941804,\"first_name\":\"Korney\",\"last_name\":\"Apushov\",\"city\":166,\"user_id\":2941804},{\"uid\":3109123,\"first_name\":\"Andrey\",\"last_name\":\"Volkov\",\"city\":4825349,\"user_id\":3109123},{\"uid\":3304214,\"first_name\":\"Bosya\",\"last_name\":\"Buvaeva\",\"city\":2,\"user_id\":3304214},{\"uid\":3777491,\"first_name\":\"Tanya\",\"last_name\":\"Elmanova\",\"city\":1,\"user_id\":3777491},{\"uid\":3820570,\"first_name\":\"Yuska\",\"last_name\":\"Pchyolkina\",\"deactivated\":\"banned\",\"user_id\":3820570},{\"uid\":3841210,\"first_name\":\"Mingian\",\"last_name\":\"Tserenov\",\"bdate\":\"6.11\",\"city\":166,\"user_id\":3841210},{\"uid\":4011096,\"first_name\":\"Kotinov\",\"last_name\":\"Badma-Khalgaevich\",\"city\":1,\"user_id\":4011096},{\"uid\":4286982,\"first_name\":\"Alexey\",\"last_name\":\"Borisov\",\"bdate\":\"19.8.1984\",\"city\":141,\"user_id\":4286982},{\"uid\":4334865,\"first_name\":\"Andrey\",\"last_name\":\"Eliseev\",\"deactivated\":\"banned\",\"user_id\":4334865},{\"uid\":4355658,\"first_name\":\"Nimgir\",\"last_name\":\"Bochkaev\",\"bdate\":\"22.7.1980\",\"city\":0,\"user_id\":4355658},{\"uid\":4543612,\"first_name\":\"Yulia\",\"last_name\":\"Bochkaeva\",\"bdate\":\"28.2\",\"city\":1,\"user_id\":4543612},{\"uid\":4613595,\"first_name\":\"Valeria\",\"last_name\":\"Golovina\",\"city\":1,\"user_id\":4613595},{\"uid\":4614686,\"first_name\":\"Andrey\",\"last_name\":\"Suchkov\",\"bdate\":\"3.11.1992\",\"city\":1,\"user_id\":4614686},{\"uid\":4735475,\"first_name\":\"Ksyusha\",\"last_name\":\"Safonova\",\"bdate\":\"7.9\",\"city\":1,\"user_id\":4735475},{\"uid\":4794458,\"first_name\":\"Ays\",\"last_name\":\"Bembeev\",\"city\":0,\"user_id\":4794458},{\"uid\":4816227,\"first_name\":\"Grigory\",\"last_name\":\"Pomadchin\",\"bdate\":\"17.4\",\"city\":1,\"user_id\":4816227},{\"uid\":5100480,\"first_name\":\"Mishka\",\"last_name\":\"Ermak\",\"bdate\":\"15.10.1993\",\"city\":1,\"user_id\":5100480},{\"uid\":5214703,\"first_name\":\"Igor\",\"last_name\":\"Bushmelev\",\"bdate\":\"13.11.1992\",\"city\":1,\"user_id\":5214703},{\"uid\":5294930,\"first_name\":\"Dmitry\",\"last_name\":\"Kostyaev\",\"bdate\":\"13.4\",\"city\":1,\"user_id\":5294930},{\"uid\":5633395,\"first_name\":\"Pavel\",\"last_name\":\"Ivanov\",\"city\":1,\"user_id\":5633395},{\"uid\":5883719,\"first_name\":\"Gelya\",\"last_name\":\"Bambueva\",\"city\":1,\"user_id\":5883719},{\"uid\":6097145,\"first_name\":\"Katya\",\"last_name\":\"Masgutova\",\"bdate\":\"3.3\",\"city\":1,\"user_id\":6097145},{\"uid\":6670989,\"first_name\":\"Anya\",\"last_name\":\"Klimova\",\"bdate\":\"14.12.1991\",\"city\":1,\"user_id\":6670989}]}"
	json.Unmarshal([]byte(resp.String()), &friends)
	return friends
}

func urlFriends(userid int, count int) string {
	return fmt.Sprintf("https://api.vk.com/method/friends.get?fields=bdate,city,sex,id&access_token=%v&v=V&user_id=%v&count=%v&v=%v", getToken(), userid, count, VERSION)
}

func chartInfo(h Human) Cchartinfo {
	cityName := h.City.Title
	if cityName == "" {
		return Cchartinfo{}
	}
	lat, lon := geocode(cityName)
	h.Lat = lat
	h.Lon = lon
	var chartInfo Cchartinfo
	if lat == 0 {
		return chartInfo
	}
	resp, _ := resty.R().Get(chartUrl(h))
	//resp, err := chartReq(h)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println(resp)
	//resp, _ := ioutil.ReadFile("./chartinfo.xml")
	xml.Unmarshal([]byte(resp.String()), &chartInfo)
	return chartInfo
}

func chartUrl(h Human) string {
	day, month, year := parseBdate(h.Bdate)
	return fmt.Sprintf("http://localhost:8080/chartinfo?name=%v&city=%v&lat=%v&lon=%v&year=%v&month=%v&day=%v&time=%v&hsys=P&display=0,1,14,4,23,10,2,3,5,6,9,7,8",
		h.FirstName + "%20" + h.LastName, h.City.ID, h.Lat, h.Lon, year, month, day, "7.5")

	//opt := []string{
	//	"-edir/home/gg/projects/go/pkg/mod/github.com/!destiny!lab/go-swetest@v0.2.0/resources/",
	//	"-b11.11.2017",
	//	"-ut00:00:00",
	//	"-geopos"+ fmt.Sprintf("%f", h.Lat) + "," + fmt.Sprintf("%f", h.Lon),
	//	"-hsysP" ,
	//	//"-a",
	//	//"-fPLBRS",
	//	//"-eswe",
	//	//"-head",
	//}
	//s := swetest.New()
	//res, err := s.Query(opt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//
	//return string(res), nil
}

func parseBdate(bdate string) (uint64, uint64, uint64) {
	s := strings.Split(bdate, ".")
	if s == nil || s[0] == "" || s[1] == "" || s[2] == "" {
		return 0, 0, 0
	}
	day, _ := strconv.ParseUint(s[0], 10, 32)
	month, _ := strconv.ParseUint(s[1], 10, 32)
	year, _ := strconv.ParseUint(s[2], 10, 32)
	return day, month, year
}

func geocode(city string) (float64, float64) {
	//fmt.Println(city)
	resp, _ := resty.R().Get(urlGeocode(city))
	//fmt.Println(resp)
	var respObj GeocodeResponse
	json.Unmarshal([]byte(resp.String()), &respObj)
	featureMembers := respObj.Response.GeoObjectCollection.FeatureMember
	var latLonStr string
	if len(featureMembers) > 1 {
		latLonStr = featureMembers[0].GeoObject.Point.Pos
	} else {
		return 0, 0
	}
	if latLonStr == "" {
		return 0, 0
	}
	pairCo := strings.Split(latLonStr, " ")
	lat, _ := strconv.ParseFloat(pairCo[1], 64)
	lon, _ := strconv.ParseFloat(pairCo[0], 64)
	return lat, lon
}

func urlGeocode(address string) string {
	return fmt.Sprintf("https://geocode-maps.yandex.ru/1.x?geocode=%v&apikey=37a46563-f5d8-4398-b116-73b788b440f1&format=json", address)
}
