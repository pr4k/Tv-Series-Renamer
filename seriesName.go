package main

import (
	"regexp"
	"strings"
)

func seriesName(Name string) string {
	re := regexp.MustCompile("[0-9]+")

	series := re.Split(Name, -1)[0][:len(re.Split(Name, -1)[0])-1]

	//fmt.Println(series)
	temp := []string{}
	final := []string{}

	temp = strings.Split(series, ".")
	final = append(final, strings.Join(temp[:], " "))

	//fmt.Println(final)
	//fmt.Println(series)
	return final[0][:len(final[0])-1]
}
