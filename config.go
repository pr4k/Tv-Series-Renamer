package main

import (
	"fmt"
	"strings"
)

type url struct {
}

const apikey = "<Enter-your-apikey>"
const username = "<Enter-your-username>"
const userkey = "<Enter-your-userKey>"

func (u url) returnSeriesID(name string) string {
	namearray := strings.Split(name, " ")
	url := fmt.Sprintf("https://api.thetvdb.com/search/series?name=%s", strings.Join(namearray[:], "%20")+"%20")
	return url
}

func (u url) returnEpisodes(seriesID int) string {
	url := fmt.Sprintf("https://api.thetvdb.com/series/%d/episodes", seriesID)
	return url
}
func (u url) returnEpisodeByPage(seriesID int, index int) string {
	url := fmt.Sprintf("https://api.thetvdb.com/series/%d/episodes?page=%d", seriesID, index)
	return url

}

func getCredentials() string {
	return fmt.Sprintf(`{"apikey":"%s","username":"%s","userkey":"%s"}`, apikey, username, userkey)
}
