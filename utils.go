package main

import (
	"os"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// IsNotDirectory : Returns true if path is not of a directory
func IsNotDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return true, err
	}
	if fileInfo.IsDir() {
		return false, err
	}
	return true, err

}
