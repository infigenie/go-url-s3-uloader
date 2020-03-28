package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/infigenie/go-url-s3-uloader/downloader"
)

// const WORKER_COUNT = 10
// const JOB_COUNT = 10

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

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
		wg.Add(1)
		// call the csv processing
		go downloader.ProcessCsvFile(nil, &wg, file)
	}
	wg.Wait()
	fmt.Println(downloader.Gray, "Kudos...!! All Processing done.")

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
