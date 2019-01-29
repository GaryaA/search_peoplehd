package main

import (
	"fmt"
	"github.com/go-resty/resty"
	"strings"
	"time"
)

type FriendsResponse struct {
	Body []Human `json:"response"`
}

type Human struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bdate     string `json:"bdate"`
	CityId    int    `json:"city"`
	Sex       int    `json:"sex"`
	Lat       float64
	Lon       float64
}

type CitiesResponse struct {
	Body struct {
		Count int    `json:"count"`
		Items []City `json:"items"`
	} `json:"response"`
}

type City struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Important int    `json:"important"`
}

type CityResponse struct {
	Body []struct {
		Cid  int    `json:"cid"`
		Name string `json:"name"`
	} `json:"response"`
}

type GeocodeResponse struct {
	Response struct {
		GeoObjectCollection struct {
			FeatureMember []struct {
				GeoObject struct {
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

const UserId = 10144987
const ClientId = 6785534

func main() {

	startUpToken()
	startUpCities()
	goscan(UserId, "ROOT", false, 1000)
}

func goscan(userid int, username string, stop bool, count int) {
	friends := friendsInfo(userid, count)
	for _, v := range friends.Body {
		if v.CityId == 0 || strings.Count(v.Bdate, ".") < 2 || v.Sex == 2 {
			continue
		}
		ci := chartInfo(v)

		if check1(ci) {
			fmt.Println(username + "                                          " + v.FirstName + " " + v.LastName + "                 " + v.Bdate)
			//fmt.Println(ci)
			fmt.Println("------------------------------------------------------------------------------------------------")
		} else {
			//fmt.Println("------------------------------------------------------------------------------------------------")
		}
		time.Sleep(20 * time.Millisecond)
		if !stop {
			goscan(v.UserID, v.FirstName+" "+v.LastName, true, 100)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check1(c Cchartinfo) bool {
	if c == (Cchartinfo{}) || c.Caspects == nil {
		return false
	}
	//for _, v := range c.Caspects.CConjunction {
	//	if (v.Attrbody1 == "Mercury" && v.Attrbody2 == "Mars") || (v.Attrbody2 == "Mercury" && v.Attrbody1 == "Mars") {
	//		return true
	//	}
	//}
	//if c.Cbodies.CSun.Attrsign_name == "Leo" && c.Cbodies.CMars.Attrsign_name == "Cancer" {
	if c.Cbodies.CVenus.Attrsign_name == "Virgo" {
		return true
	}
	return false
}

func checkResp(r *resty.Response) {
	message := string(r.String())
	if strings.Contains(message, "\"error_description\"") {
		panic(message)
	}
}
