package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Product struct {
	Token string
}
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
	allseries  []episode `json: "allseries"`
	SeasonName string
}
type seasoninfo struct {
	Aliases    []string
	Banner     string
	FirstAired string
	Id         int
	Network    string
	Overview   string
	SeriesName string
	Slug       string
	Status     string
}

func main() {

	tk := token()
	//
	//fmt.Println(similar(name))
	path := os.Args[1:][0]
	name := files(path)
	fmt.Println(name)
	var allepisodesdata []database
	var allepisodes []episode
	seriesname := similar(name)
	fmt.Println(seriesname)
	for _, j := range seriesname {
		newpath := path + "/" + j
		if _, err := os.Stat(newpath); os.IsNotExist(err) {
			os.Mkdir(newpath, os.ModePerm)

		}
	}
	for _, name := range seriesname {
		series := seriesid(name, tk)
		for _, i := range series.Data {
			fmt.Println(i.Id)
		}

		allepisodes = episodes(series.Data[0].Id, tk, series.Data[0].SeriesName)
		allepisodesdata = append(allepisodesdata, database{allepisodes, series.Data[0].SeriesName})
	}

	rename(name, path, allepisodesdata)
}

func files(path string) []string {
	file, err := ioutil.ReadDir(path)
	var names []string
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(file)
	for i, f := range file {
		IsNotFolder, _ := IsNotDirectory(path + "/" + f.Name())
		if len(strings.Split(f.Name(), ".")) != 1 && IsNotFolder && strings.Split(f.Name(), ".")[len(strings.Split(f.Name(), "."))-1] != "txt" && strings.Split(f.Name(), ".")[len(strings.Split(f.Name(), "."))-1] != "html" {
			//if strings.Split(f.Name(),".")[len(strings.Split(f.Name(),"."))-1]!="srt"{
			names = append(names, f.Name())
			fmt.Println(i, f.Name())
			//}
		}
	}
	//fmt.Println(names, len(names))
	return names
}

func similar(names []string) []string {
	series := []string{}

	for i, _ := range names {
		if contains(series, seriesName(names[i])) == false {

			series = append(series, seriesName(names[i]))
		}
	}
	//fmt.Println(series)

	//fmt.Println(final)
	fmt.Println(series)
	return series
	//return final
}
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func seriesid(name string, token string) season {
	namearray := strings.Split(name, " ")
	url := fmt.Sprintf("https://api.thetvdb.com/search/series?name=%s", strings.Join(namearray[:], "%20")+"%20")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s := string(body)
	//fmt.Println(s)
	var series season
	err := json.Unmarshal([]byte(s), &series)
	if err == nil {
		//		fmt.Println(series,"yeey")
	} else {
		fmt.Println("wrong")
	}
	return series

}

func episodes(seriesid int, token string, SeriesName string) []episode {
	url := fmt.Sprintf("https://api.thetvdb.com/series/%d/episodes", seriesid)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s := string(body)
	//fmt.Println(s)oldname
	var result map[string]interface{}
	json.Unmarshal([]byte(s), &result)
	allepisodes := make([]episode, 0)
	fmt.Println("Total pages", result["links"].(map[string]interface{})["last"])
	for _, j := range result["data"].([]interface{}) {
		name := j.(map[string]interface{})["episodeName"].(string)
		season := j.(map[string]interface{})["airedSeason"].(float64)
		episodenumber := j.(map[string]interface{})["airedEpisodeNumber"].(float64)
		allepisodes = append(allepisodes, episode{season, episodenumber, name, SeriesName})
	}

	if int(result["links"].(map[string]interface{})["last"].(float64)) > 1 {
		for index := 2; index <= int(result["links"].(map[string]interface{})["last"].(float64)); index++ {
			url := fmt.Sprintf("https://api.thetvdb.com/series/%d/episodes?page=%d", seriesid, index)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			res, _ := http.DefaultClient.Do(req)
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
				allepisodes = append(allepisodes, episode{season, episodenumber, name, seriesName(SeriesName)})
			}
		}

	}
	//fmt.Println(allepisodes)
	return allepisodes
}
func rename(names []string, path string, allepisodesdata []database) {
	oldname := make([]episode, 0)
	re := regexp.MustCompile("([a-zA-Z]+)")
	for _, name := range names {
		txt := "Have9834a908123great10891819081day!"
		txt = "How.I.Met.Your.Mother.S01E01.1080p.WEB.DL.Farda.DL.mkv"
		txt = name
		fmt.Println(txt, 1)
		counter := 0
		split := re.Split(txt, -1)
		//fmt.Println(split, len(split))
		var data []int

		for _, j := range split {
			for _, k := range strings.Split(j, " ") {
				a, err := strconv.Atoi(k)
				if err == nil {
					if a > 99 {
						data = append(data, a/100)
						data = append(data, a%100)
						counter = 2
						break
					} else {
						data = append(data, a)
						counter = counter + 1
						if counter >= 2 {
							break
						}
					}
				} else {
					b := strings.Split(k, ".")

					for _, i := range b {
						a, err := strconv.Atoi(i)
						if err == nil {
							data = append(data, a)
							counter = counter + 1
							if counter >= 2 {
								break
							}
							fmt.Println(b, k, data, i)

						}

					}
					b = strings.Split(k, "_")
					for _, i := range b {
						a, err := strconv.Atoi(i)
						if err == nil {
							data = append(data, a)
							counter = counter + 1
							if counter >= 2 {
								break
							}
							fmt.Println(b, k, data, i)

						}

					}
					//fmt.Println(data)
					if counter >= 2 {
						break
					}
				}
			}
			if counter >= 2 {
				break
			}

		}
		fmt.Println(seriesName(name), data)
		oldname = append(oldname, episode{float64(data[0]), float64(data[1]), name, seriesName(name)})

	}

	var finalpath string
	for _, j := range oldname {

		//if exists(path+fmt.Sprintf("s%d",int(j.season))){
		for _, episodedata := range allepisodesdata {

			if j.SeriesName == episodedata.SeasonName {

				var temppath string
				temppath = path + "/" + j.SeriesName + "/" + fmt.Sprintf("/S%d", int(oldname[0].season))
				for _, j := range oldname {
					var newpath string
					if int(j.season) > 9 {
						newpath = path + "/" + j.SeriesName + "/" + fmt.Sprintf("/S%d", int(j.season))
					} else {
						newpath = path + "/" + j.SeriesName + "/" + fmt.Sprintf("/S0%d", int(j.season))
					}
					if temppath != newpath {
						if _, err := os.Stat(newpath); os.IsNotExist(err) {
							os.Mkdir(newpath, os.ModePerm)
						}
						temppath = newpath
					}

				}

				for _, k := range episodedata.allseries {
					if k.season == j.season && k.episodenumber == j.episodenumber {
						ext := strings.Split(j.name, ".")[len(strings.Split(j.name, "."))-1]
						if int(j.season) > 9 {
							if int(j.episodenumber) > 9 {
								//fmt.Println(path+"/"+j.name, path+fmt.Sprintf("/S%d/E%d-%s", int(j.season), int(j.episodenumber), k.name))
								finalpath = fmt.Sprintf("/S%d/E%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext)
								//os.Rename(path+"/"+j.name, path+fmt.Sprintf("/S%d/E%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext))
							} else {
								//fmt.Println(path+"/"+j.name, path+fmt.Sprintf("/S%d/E0%d-%s", int(j.season), int(j.episodenumber), k.name))
								finalpath = fmt.Sprintf("/S%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext)
								//os.Rename(path+"/"+j.name, path+fmt.Sprintf("/S%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext))
							}
						} else {
							if int(j.episodenumber) > 9 {
								//temp:=strings.Split(k.name,"?")
								//l:=strings.Join(temp[:], "")
								l := k.name
								//fmt.Println(path+j.name, path+fmt.Sprintf("/S0%d/E%d-%s", int(j.season), int(j.episodenumber), l))
								finalpath = fmt.Sprintf("/S0%d/E%d-%s.%s", int(j.season), int(j.episodenumber), l, ext)
								//os.Rename(path+"/"+j.name, path+fmt.Sprintf("/S0%d/E%d-%s.%s", int(j.season), int(j.episodenumber), l, ext))
							} else {
								//temp:=strings.Split(k.name,"?")
								//l:=strings.Join(temp[:], "")
								l := k.name
								//fmt.Println(path+"/"+j.name, path+fmt.Sprintf("/S0%d/E0%d-%s", int(j.season), int(j.episodenumber), l))
								finalpath = fmt.Sprintf("/S0%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), l, ext)
								//os.Rename(path+"/"+j.name, path+fmt.Sprintf("/S0%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), l, ext))
							}
						}
						fmt.Println(path + "/" + j.SeriesName + "/" + finalpath)
						finalpath = path + "/" + j.SeriesName + "/" + finalpath
						os.Rename(path+"/"+j.name, finalpath)
						//}
					}
					//os.Rename(path,"/media/pr4k/New Volume/arrow/s02/Arrow.S02E02.1080p.BluRay-[Bi-3-Seda.Ir].mkv.mkv")
				}
			}
		}
	}
}
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
func IsNotDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return true, err
	}
	if fileInfo.IsDir() {
		return false, err
	} else {
		return true, err
	}
}
