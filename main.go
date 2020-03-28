package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/infigenie/go-url-s3-uloader/downloader"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// const WORKER_COUNT = 10
// const JOB_COUNT = 10

func main() {
	log.Println("starting application...")
	var files []string

	root := "./"
	err := filepath.Walk(root, Visit(&files))
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
	for _, file := range files {
		fmt.Println("Processing file: " + file)
		// call the csv processing
		downloader.ProcessCsvFile(nil, file)
	}

}

// Visit visits the folder
func Visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if filepath.Ext(path) != ".csv" {
			return nil
		}
		*files = append(*files, path)
		return nil
	}
}
