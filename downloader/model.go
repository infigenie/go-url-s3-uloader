package downloader

// URL represents the structure for storing GTINs in gtins table in DB.
type URL struct {
	//ID          int
	Address        string `csv:"0"`
	outputFilename string
}

// Job defines a GTIN generation JOB
type Job struct {
	ID           int
	Start        int
	End          int
	URLs         []URL
	urlsFilename string
}

// Result represents a result of GS1 GTIN Check Job.
type Result struct {
	job      Job
	response URLResponse
}

// URLResponse represents a result of insertion of gtins into gtins table in DB.
type URLResponse struct {
	Job  Job    `json:"job,omitempty"`
	Resp string `json:"response,omitempty"`
	Err  string `json:"error,omitempty"`
}

// GenResponse represetns the aggregated list of DBResposes.
type GenResponse struct {
	Res *[]GenResponse `json:"success,omitempty"`
}
