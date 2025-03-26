package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

// # usage: ./crawler URL maxConcurrency maxPages
func main() {
	//check and get the argument
	base_url, maxConcurrency, maxPages := process_arguments(os.Args)
	//parse the URL
	parsedUrl, err := url.Parse(base_url)
	//if error
	if err != nil {
		//log error
		fmt.Println(err)
		//stop
		os.Exit(1)
	}

	//set the config to be used by the go routines when crawling
	cnf := config{
		//maximum amount of pages to be checked
		maxPages: maxPages,
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

	//check if limit was reached
	if cnf.page_limit_reached() {
		fmt.Println("Page limit reached")
	}

	// finish message
	fmt.Println("Finished crawl, time elapsed: ", time.Since(time_start))

	//for each key value pair (page found)
	for key, value := range cnf.pages {
		//print information
		fmt.Println(value, "x ", key)
	}
}

// function that checks if there are enough arguments and returns the normalized url
func process_arguments(arguments []string) (string, int, int) {
	// check if we have enough arguments (should be more than 1)
	if len(arguments) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	// check if we have 3 argument
	if len(arguments) > 4 {
		fmt.Println("too many arguments provided (need 3, have", len(arguments), ")")
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

	//set default values
	maxConcurrency := 1
	max_pages := 100

	//if more than URL provided
	if len(arguments) > 2 {
		// string to int
		maxConcurrency, err = strconv.Atoi(arguments[2])
		if err != nil {
			//log error
			fmt.Println(err)
			//stop
			os.Exit(1)
		}
		//failsafe
		if maxConcurrency < 1 {
			maxConcurrency = 1
		}
		//if more than concurrency limit provided
		if len(arguments) > 3 {
			// string to int
			max_pages, err = strconv.Atoi(arguments[3])
			if err != nil {
				//log error
				fmt.Println(err)
				//stop
				os.Exit(1)
			}
			//failsafe
			if max_pages < 1 {
				max_pages = 100
			}
		}
	}
	fmt.Println("Config: ", base_url, maxConcurrency, max_pages)
	//return the arguments
	return base_url, maxConcurrency, max_pages
}

/*
https://gobyexample.com/command-line-arguments
go build -o crawler && ./crawler BASE_URL
go build -o crawler && ./crawler google.nl
go build -o crawler && ./crawler https://wagslane.dev/
*/
