package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"
)

func main() {
	//check and get the argument
	base_url := process_arguments(os.Args)
	//parse the URL
	parsedUrl, err := url.Parse(base_url)
	//if error
	if err != nil {
		//log error
		fmt.Println(err)
		//stop
		os.Exit(1)
	}

	//config
	//set max concurrency of workers
	maxConcurrency := 0

	//failsafe
	if maxConcurrency < 1 {
		maxConcurrency = 1
	}
	
	//set the config to be used by the go routines when crawling
	cnf := config{
		//empty map, base url is added
		pages: map[string]int{},
		//save the base url
		baseURL: parsedUrl,
		//create wait group
		wg: new(sync.WaitGroup),
		//create buffer for limiting go routines
		concurrencyControl: make(chan struct{}, maxConcurrency),
		//create mutex,
		mu: new(sync.Mutex),
	}
	
	//get start time
	time_start := time.Now()
	//start message
	fmt.Printf("Starting crawl of: %v\n", base_url)
	
	//start crawling at the supplied url
	//add the base url
	cnf.addPageVisit(base_url)
	
	//this should spawn more workers, to a max of x [maxConcurrency] workers
	cnf.crawl_page(base_url)
	
	//wait for all routines to end
	cnf.wg.Wait()
	
	// finish message
	fmt.Println("Finished crawl, time elapsed: ", time.Now().Sub(time_start))
	
	//for each key value pair (page found)
	for key, value := range cnf.pages {
		//print information
		fmt.Println(value, "x ", key)
	}
}

// function that checks if there are enough arguments and returns the normalized url
func process_arguments(arguments []string) string {
	// check if we have enough arguments (should be more than 1)
	if len(arguments) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	// check if we only have 1 argument
	if len(arguments) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	// get the url
	BASE_URL := arguments[1]

	// normalize URL
	base_url, err := normalizeURL(BASE_URL)
	// if error
	if err != nil {
		//log error
		fmt.Println(err)
		//stop
		os.Exit(1)
	}
	//return the url string
	return base_url
}

/*
https://gobyexample.com/command-line-arguments
go build -o crawler && ./crawler BASE_URL
go build -o crawler && ./crawler google.nl
go build -o crawler && ./crawler https://wagslane.dev/
*/
