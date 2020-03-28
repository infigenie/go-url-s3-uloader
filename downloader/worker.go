package downloader

// import (
// 	"context"
// 	"strconv"
// )

// // calculateCheckDigit Calculates and returns a check digit for a given GTIN.
// func calculateCheckDigit(gtin string) (int, error) {

// 	// Take an array of multipliers.
// 	multipliers := [17]int{3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3}

// 	var sum int
// 	for i := 0; i < len(gtin); i++ {
// 		v, err := strconv.Atoi(string(gtin[i]))
// 		if err != nil {
// 			return sum, err
// 		}
// 		sum = sum + v*multipliers[len(multipliers)-len(gtin)+i]
// 	}

// 	var round int
// 	if round = sum; sum%10 != 0 {
// 		round = (10 - sum%10) + sum
// 	}

// 	return round - sum, nil
// }

// func doGeneration(ctx context.Context, gcp string) ([]DBResponse, error) {

// 	var s, e string
// 	// make the gcp to pad zeros
// 	for i := 0; i < 13-len(gcp)-1; i++ {
// 		s = s + "0"
// 		e = e + "9"
// 	}

// 	sgcp := gcp + s
// 	egcp := gcp + e
// 	start, _ := strconv.Atoi(sgcp)
// 	end, _ := strconv.Atoi(egcp)

// 	noOfGtins, _ := strconv.Atoi("1" + s)
// 	batchSize := 10000
// 	noOfJobs := noOfGtins / batchSize
// 	if noOfGtins < batchSize {
// 		noOfJobs = 1
// 		batchSize = noOfGtins
// 	}

// 	var jobs = make(chan Job, noOfJobs)
// 	var results = make(chan Result, noOfJobs)

// 	res := []DBResponse{}

// 	go allocate(noOfJobs, gcp, start, end, batchSize, &jobs)
// 	done := make(chan bool)
// 	go result(done, &res, &results)

// 	// Bound noOfWorkers(Go routines) making call to GS1 Cloud
// 	noOfWorkers := noOfJobs
// 	if noOfJobs >= 10 {
// 		noOfWorkers = 10
// 	}

// 	createWorkerPool(ctx, noOfWorkers, &results, &jobs)
// 	<-done
// 	//fmt.Println(res)
// 	return res, nil
// }
