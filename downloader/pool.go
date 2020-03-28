package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

// worker is a worker process for  and processing GS1 Check API to validate a GTIN.
func worker(ctx context.Context, wg *sync.WaitGroup, results *chan Result, jobs *chan Job) {
	for job := range *jobs {
		// change here later for batches...
		for i := range job.URLs {
			output := Result{job, DownloadFile(job, i)}
			*results <- output
		}
	}
	wg.Done()
}

// createWorkerPool creates a pool of workers for calling GS1 Check API to validate a number of GTINs.
func createWorkerPool(ctx context.Context, noOfWorkers int, results *chan Result, jobs *chan Job) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, results, jobs)
	}
	wg.Wait()
	close(*results)
}

// allocate  creates jobs by assigning a Job id and and gtin.
func allocate(urlsFilename string, noOfJobs int, urls []URL, startIndex int, batchSize int, jobs *chan Job) {
	for i := 0; i < noOfJobs; i++ {
		s := startIndex + i*batchSize
		e := s + batchSize
		job := Job{i, s, e, urls[s:e], urlsFilename}
		fmt.Println(Yellow, "Allocating Job: ", job.ID, job.URLs, job.urlsFilename)
		*jobs <- job
	}
	close(*jobs)
}

// result process the reads from the result channel and finally closes everything by closing the done channel(Done=true)
func result(done chan bool, res *[]URLResponse, results *chan Result) {
	for result := range *results {
		*res = append(*res, result.response)
	}
	done <- true
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(job Job, loc int) URLResponse {
	fmt.Println(Cyan, "Processing job:", job.ID, job.URLs, job.urlsFilename)
	// Get the data
	resp, err := http.Get(job.URLs[loc].Address)
	if err != nil {
		return URLResponse{Job: job, Resp: "", Err: err.Error()}
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(job.URLs[loc].outputFilename)
	if err != nil {
		return URLResponse{Job: job, Resp: "", Err: err.Error()}
	}
	fmt.Println(Green, "Saving job:", job.ID, job.urlsFilename, job.URLs[loc].outputFilename)

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return URLResponse{Job: job, Resp: "Success", Err: ""}
}
