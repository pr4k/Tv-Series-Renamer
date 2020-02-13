package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type episode struct {
	season        float64
	episodenumber float64
	name          string
	SeriesName    string
}
type season struct {
	Data []seasoninfo `json:"data"`
}
type database struct {
	Allseries  []episode `json:"allseries"`
	SeasonName string
}
type seasoninfo struct {
	Aliases    []string
	Banner     string
	FirstAired string
	ID         int
	Network    string
	Overview   string
	SeriesName string
	Slug       string
	Status     string
}

func getSeriesNames(names []string) []string {
	series := []string{}
	for i := range names {
		tempName := seriesNameRegex(names[i])
		if contains(series, tempName) == false {

			series = append(series, tempName)
		}
	}
	return series

}
func seriesNameRegex(Name string) string {
	re := regexp.MustCompile("[0-9]+")
	series := re.Split(Name, -1)[0][:len(re.Split(Name, -1)[0])-1]
	var temp []string

	temp = strings.Split(series, ".")

	return strings.TrimSpace(strings.Join(temp, " "))
}

func generateDatabase(seriesname []string, tk string) ([]database, error) {
	var geturl url
	var allepisodes []episode
	var allepisodesdata []database
	for _, name := range seriesname {

		series, err := seriesid(name, tk, geturl)
		if err != nil {
			return nil, err
		}

		allepisodes = episodes(series.Data[0].ID, tk, series.Data[0].SeriesName, geturl)
		allepisodesdata = append(allepisodesdata, database{allepisodes, series.Data[0].SeriesName})
	}
	return allepisodesdata, nil
}
func seriesid(name string, token string, geturl url) (season, error) {
	var series season
	url := geturl.returnSeriesID(name)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {

		return series, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s := string(body)

	err = json.Unmarshal([]byte(s), &series)
	if err != nil {
		return series, err
	}
	return series, nil

}

func episodes(seriesid int, token string, SeriesName string, geturl url) []episode {
	url := geturl.returnEpisodes(seriesid)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s := string(body)
	var result map[string]interface{}
	json.Unmarshal([]byte(s), &result)
	allepisodes := make([]episode, 0)
	for _, j := range result["data"].([]interface{}) {
		name := j.(map[string]interface{})["episodeName"].(string)
		season := j.(map[string]interface{})["airedSeason"].(float64)
		episodenumber := j.(map[string]interface{})["airedEpisodeNumber"].(float64)
		allepisodes = append(allepisodes, episode{season, episodenumber, name, SeriesName})
	}

	if int(result["links"].(map[string]interface{})["last"].(float64)) > 1 {
		for index := 2; index <= int(result["links"].(map[string]interface{})["last"].(float64)); index++ {
			url := geturl.returnEpisodeByPage(seriesid, index)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			s := string(body)
			//fmt.Println(s)
			var result map[string]interface{}
			json.Unmarshal([]byte(s), &result)
			for _, j := range result["data"].([]interface{}) {
				name := j.(map[string]interface{})["episodeName"].(string)
				season := j.(map[string]interface{})["airedSeason"].(float64)
				episodenumber := j.(map[string]interface{})["airedEpisodeNumber"].(float64)
				allepisodes = append(allepisodes, episode{season, episodenumber, name, seriesNameRegex(SeriesName)})
			}
		}

	}

	return allepisodes
}
