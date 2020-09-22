package session

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FindByGithub(token string) []byte {
	to := "https://api.github.com/user"
	req, _ := http.NewRequest("GET", to, nil)
	req.Header.Set("Authorization", token)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	return byteArray
}
