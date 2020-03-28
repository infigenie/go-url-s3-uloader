package downloader

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gocarina/gocsv"
)

// ProcessCsvFile to process the csv
func ProcessCsvFile(ctx context.Context, wg *sync.WaitGroup, fileName string) error {

	defer wg.Done()

	urlsFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer urlsFile.Close()

	var urlsFromFile []URL
	var urls []URL

	if err := gocsv.UnmarshalFile(urlsFile, &urlsFromFile); err != nil { // Load clients from file
		panic(err)
	}

	folderName := strings.Split(fileName, ".")[0]
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, os.ModePerm)
	}
	// do the worrk from here...
	for _, url := range urlsFromFile {
		if url.Address == "" {
			continue
		}

		fmt.Println(fileName+" : ", url)
		fnArray := strings.Split(url.Address, "/")
		outputFilenName := folderName + "/" + fnArray[len(fnArray)-1]
		urls = append(urls, URL{Address: url.Address, outputFilename: outputFilenName})
	}
	fmt.Println(len(urls))
	DownloadAndSaveURLs(ctx, fileName, urls)
	fmt.Println(Red, "All Urls saved for: ", fileName)
	return nil
}

// DownloadAndSaveURLs removes a GCP from the database.
func DownloadAndSaveURLs(ctx context.Context, urlsFileName string, urls []URL) ([]URLResponse, error) {

	noOfURLs := len(urls)
	var jobs = make(chan Job, noOfURLs)
	var results = make(chan Result, noOfURLs)
	start := 0
	batchSize := 1
	res := []URLResponse{}
	go allocate(urlsFileName, noOfURLs, urls, start, batchSize, &jobs)
	done := make(chan bool)
	go result(done, &res, &results)

	// Bound noOfWorkers(Go routines) making call to GS1 Cloud
	noOfWorkers := noOfURLs
	// if noOfURLs >= 100 {
	// 	noOfWorkers = 100
	// }

	createWorkerPool(ctx, noOfWorkers, &results, &jobs)
	<-done

	return res, nil
}
