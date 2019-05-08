package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
)
func token() string {
	var tokenkey Product
	url := "https://api.thetvdb.com/login"

	var apikeys = []byte(`{"apikey":"Your_Api_key","username":"Your_Username","userkey":"Your_User_Key"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(apikeys))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err:= http.DefaultClient.Do(req)
	if err !=nil{
		fmt.Println(err)
		return "error"
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s := string(body)
	err = json.Unmarshal([]byte(s), &tokenkey)
	if err == nil {
		//		fmt.Println(tokenkey)
	} else {
		fmt.Println(err)
		//		fmt.Printf("new")
	}
	key := tokenkey.Token
	return key
}