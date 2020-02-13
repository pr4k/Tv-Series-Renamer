package main

import (
	"log"
	"os"

	"github.com/akamensky/argparse"
)

func main() {

	parser := argparse.NewParser("Tv Series Renamer", "Organizes your tvseries")
	inputPath := parser.String("i", "inputPath", &argparse.Options{Required: true, Help: "path/to/folder/for/input"})

	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	tk := token()

	name := files(*inputPath)
	var allepisodesdata []database
	seriesname := getSeriesNames(name)
	createDirectory(seriesname, *inputPath)
	allepisodesdata, err = generateDatabase(seriesname, tk)
	if err != nil {
		log.Fatal(err)
	}

	rename(name, *inputPath, allepisodesdata)

}
