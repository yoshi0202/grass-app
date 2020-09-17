package services

import (
	"io/ioutil"
	"net/http"
)

func Get(param string) []byte {
	url := "https://github.com/users/" + param + "/contributions"
	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray
}
