package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"net/url"
)

const root string = "https://challenge.curbside.com/"

func get_api(api string)(string) {
	url, _ := url.Parse(root)
	url.Path = path.Join(url.Path, api)
	return url.String()
}

func get_raw_response(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func get_session_id() string {
	session_api := get_api("get-session")
	apiResp := get_raw_response(session_api)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(apiResp), &data)
	if err != nil {
		panic(err)
	}
	return data["session"].(string)
}

func main() {
	fmt.Println(get_session_id())
}
