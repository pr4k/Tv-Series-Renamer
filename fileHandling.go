package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func files(path string) []string {
	file, err := ioutil.ReadDir(path)
	var names []string
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range file {
		IsNotFolder, _ := IsNotDirectory(f.Name())
		if len(strings.Split(f.Name(), ".")) != 1 && IsNotFolder && strings.Split(f.Name(), ".")[len(strings.Split(f.Name(), "."))-1] != "txt" && strings.Split(f.Name(), ".")[len(strings.Split(f.Name(), "."))-1] != "html" {
			names = append(names, f.Name())
		}
	}
	return names
}

func createDirectory(seriesname []string, path string) {
	for _, j := range seriesname {
		newpath := j
		fmt.Println(filepath.Join(path, newpath))
		if _, err := os.Stat(filepath.Join(path, newpath)); os.IsNotExist(err) {
			os.Mkdir(filepath.Join(path, newpath), os.ModePerm)

		}
	}
}
func rename(names []string, path string, allepisodesdata []database) {
	oldname := make([]episode, 0)
	re := regexp.MustCompile("([a-zA-Z]+)")
	for _, name := range names {
		txt := "Have9834a908123great10891819081day!"
		txt = "How.I.Met.Your.Mother.S01E01.1080p.WEB.DL.Farda.DL.mkv"
		txt = name

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

		oldname = append(oldname, episode{float64(data[0]), float64(data[1]), name, seriesNameRegex(name)})

	}

	var newpath string
	for _, j := range oldname {

		//if exists(path+fmt.Sprintf("s%d",int(j.season))){
		for _, episodedata := range allepisodesdata {

			if j.SeriesName == episodedata.SeasonName {

				var temppath string
				temppath = j.SeriesName + "/" + fmt.Sprintf("/S%d", int(oldname[0].season))
				for _, j := range oldname {
					var newpath string
					if int(j.season) > 9 {
						newpath = j.SeriesName + "/" + fmt.Sprintf("/S%d", int(j.season))
					} else {
						newpath = j.SeriesName + "/" + fmt.Sprintf("/S0%d", int(j.season))
					}
					if temppath != newpath {
						if _, err := os.Stat(filepath.Join(path, newpath)); os.IsNotExist(err) {
							os.Mkdir(filepath.Join(path, newpath), os.ModePerm)
						}
						temppath = newpath
					}

				}

				for _, k := range episodedata.Allseries {
					if k.season == j.season && k.episodenumber == j.episodenumber {
						ext := strings.Split(j.name, ".")[len(strings.Split(j.name, "."))-1]
						if int(j.season) > 9 {
							if int(j.episodenumber) > 9 {
								//fmt.Println(newpath.Join(path,j.name), path+fmt.Sprintf("/S%d/E%d-%s", int(j.season), int(j.episodenumber), k.name))
								newpath = fmt.Sprintf("/S%d/E%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext)
								//os.Rename(newpath.Join(path,j.name), path+fmt.Sprintf("/S%d/E%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext))
							} else {
								//fmt.Println(newpath.Join(path,j.name), path+fmt.Sprintf("/S%d/E0%d-%s", int(j.season), int(j.episodenumber), k.name))
								newpath = fmt.Sprintf("/S%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext)
								//os.Rename(newpath.Join(path,j.name), path+fmt.Sprintf("/S%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), k.name, ext))
							}
						} else {
							if int(j.episodenumber) > 9 {
								//temp:=strings.Split(k.name,"?")
								//l:=strings.Join(temp[:], "")
								l := k.name
								//fmt.Println(path+j.name, path+fmt.Sprintf("/S0%d/E%d-%s", int(j.season), int(j.episodenumber), l))
								newpath = fmt.Sprintf("/S0%d/E%d-%s.%s", int(j.season), int(j.episodenumber), l, ext)
								//os.Rename(newpath.Join(path,j.name), path+fmt.Sprintf("/S0%d/E%d-%s.%s", int(j.season), int(j.episodenumber), l, ext))
							} else {
								//temp:=strings.Split(k.name,"?")
								//l:=strings.Join(temp[:], "")
								l := k.name
								//fmt.Println(newpath.Join(path,j.name), path+fmt.Sprintf("/S0%d/E0%d-%s", int(j.season), int(j.episodenumber), l))
								newpath = fmt.Sprintf("/S0%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), l, ext)
								//os.Rename(newpath.Join(path,j.name), path+fmt.Sprintf("/S0%d/E0%d-%s.%s", int(j.season), int(j.episodenumber), l, ext))
							}
						}

						newpath := filepath.Join(path, j.SeriesName, newpath)

						fmt.Println(filepath.Join(path, j.name), " -> ", newpath)
						os.Rename(filepath.Join(path, j.name), newpath)
						//}
					}
					//os.Rename(path,"/media/pr4k/New Volume/arrow/s02/Arrow.S02E02.1080p.BluRay-[Bi-3-Seda.Ir].mkv.mkv")
				}
			}
		}

	}
}
