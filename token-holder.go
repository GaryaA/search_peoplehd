package main

import (
	"io/ioutil"
)

var vktoken string

//const AUTH_USER_ID = 650944593

func startUpToken() {
	dat, err := ioutil.ReadFile("./token")
	check(err)
	if dat == nil || string(dat) == "" {
		updateTokenRemote()
	}
	vktoken = string(dat)
}

func updateTokenRemote() {
	newToken := authVk()
	updateToken(newToken)
}

func getToken() string {
	return vktoken
}

func updateToken(newToken string) {
	d := []byte(newToken)
	err := ioutil.WriteFile("./token", d, 0644)
	check(err)
	vktoken = newToken
}

func authVk() string {

	//{
	//	resp, err := resty.
	//		R().
	//		//SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36").
	//		Get(fmt.Sprintf("https://oauth.vk.com/authorize?client_id=%v&display=page&scope=friends&redirect_uri=vk.com&response_type=code&v=5.92", CLIENT_ID))
	//	check(err)
	//	checkResp(resp)
	//	fmt.Println(resp.Status(), " ", resp.StatusCode())
	//	fmt.Println(resp)
	//}
	//{
	//	resp, err := resty.R().
	//		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36").
	//		Get(fmt.Sprintf("https://login.vk.com/?act=grant_access&client_id=%v&settings=2&redirect_uri=vk.com&response_type=code&group_ids=&token_type=0&v=5.92&state=&display=page&ip_h=cb1d40365b19e91418&hash=1546364363_a27e5dcee0a43c8c0c&https=1", CLIENT_ID))
	//	check(err)
	//	checkResp(resp)
	//	fmt.Println(resp)
	//}
	//{
	//	resp, err := resty.R().
	//		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36").
	//		Get(fmt.Sprintf("https://oauth.vk.com/access_token?client_id=%v&client_secret=RtW0DxXTo503ItUSx6RE&v=5.92&grant_type=client_credential", CLIENT_ID))
	//	check(err)
	//	checkResp(resp)
	//	return resp.String()
	//}
	panic("df")
	return ""
}
