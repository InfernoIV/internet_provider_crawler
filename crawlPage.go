package main

import (
	"fmt"
	"net/url"
	"sync"
)

// struct to be used in go routines
type config struct {
	maxPages           int
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

// new crawlPage function (using go routines)
func (cfg *config) crawl_page(rawCurrentURL string) {
	//if limit is reached
	if cfg.page_limit_reached() {
		//limit reached, stop
		return
	}
	//indication for main program to wait
	// Increment the WaitGroup counter
	cfg.wg.Add(1)
	// Decrement the counter when the goroutine completes.
	defer cfg.wg.Done()

	//ensure that not too many workers are working
	//fill the channel to indicate a worker is spawned
	cfg.concurrencyControl <- struct{}{} // Send a signal; value does not matter.
	//make sure to clear the channel when exiting
	defer cfg.defer_channel_read()

	//get normalized url
	url_current, err := normalizeURL(rawCurrentURL)
	//if error
	if err != nil {
		//log error
		fmt.Println(err)
		return
	}
	fmt.Println("Crawling", url_current)

	//get htmlBody from url (should error when connection is not possible or when content type is text/html): no need for error logging
	htmlBody, err := getHTML(url_current)
	if err != nil {
		//stop
		return
	}

	//get urls
	url_list, err := getURLsFromHTML(htmlBody, url_current)
	//if error
	if err != nil {
		//log error
		fmt.Println(err)
		//stop
		return
	}

	//for each url
	for i := 0; i < len(url_list); i++ {
		//get the url
		url_found := url_list[i]

		//get the hostname
		hostname_found, err := get_hostname(url_found)
		//if error
		if err != nil {
			//log error
			fmt.Println(err)
			//stop
			continue
		}

		//check if it is the same hostname
		if cfg.baseURL.Hostname() != hostname_found {
			//not same hostname: stop
			continue
		}

		//if seen (visited?) for the first time
		if cfg.addPageVisit(url_found) {
			//spawn routine to dig deeper
			go cfg.crawl_page(url_found)
		}
	}
}

// function that saves the page visit, and returns if it is the first time
func (cfg *config) addPageVisit(normalized_url string) (is_first bool) {
	//lock
	cfg.mu.Lock()
	//unlock
	defer cfg.mu.Unlock()

	//get the number of times this page was found (0 if not found)
	url_times_found := cfg.pages[normalized_url]
	//add the count
	cfg.pages[normalized_url] = url_times_found + 1

	//return if found for the 1st time
	return cfg.pages[normalized_url] == 1
}

// function to read channel to ensure other workers can spawn
func (cfg *config) defer_channel_read() {
	//send signal indicating stopping of worker
	<-cfg.concurrencyControl // Wait for sort to finish; discard sent value.
}

func (cfg *config) page_limit_reached() (limit_reached bool) {
	//lock
	cfg.mu.Lock()
	//unlock
	defer cfg.mu.Unlock()
	//return if we passed the limit
	return len(cfg.pages) >= cfg.maxPages
}
