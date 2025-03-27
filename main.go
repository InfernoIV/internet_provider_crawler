package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
	"sort"
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

	//print the report
	print_report(cnf.pages, base_url)
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
	base_url, err := normalize_url(BASE_URL)
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

// function that will print a report of the results
func print_report(pages map[string]int, baseURL string) {
	//struct definition	
	type url_report struct {
			url string         
			amount_found int
	} 
	
	

	//stolen from https://www.geeksforgeeks.org/how-to-sort-golang-map-by-keys-or-values/
	//make a slice
	url_slice := make([]url_report, 0, len(pages))
	for key, value := range pages {
		//add url report
		url_slice = append(url_slice, url_report{ url: key, amount_found: value})
	}

	//sort the slice (largest to smallest)
	sort.Slice(url_slice, func(i, j int) bool { return url_slice[i].amount_found > url_slice[j].amount_found })
	
	//print header
	fmt.Println("=============================")
	fmt.Println("REPORT for", baseURL)
	fmt.Println("=============================")

	//for each entry in the slice
	for _, element := range url_slice {
		//print message
		fmt.Println("Found", element.amount_found, "internal links to", element.url)
	}
}

/*
https://gobyexample.com/command-line-arguments
go build -o crawler && ./crawler BASE_URL
go build -o crawler && ./crawler google.nl
go build -o crawler && ./crawler https://wagslane.dev/
go build -o crawler && ./crawler https://wagslane.dev/ 5 100
go build -o crawler && ./crawler https://news.ycombinator.com/ 3 25
*/
