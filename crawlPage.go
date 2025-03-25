package main

import "log"

func crawlPage(raw_base_url, raw_current_url string, pages map[string]int) {
	hostname_base, err := get_hostname(raw_base_url)
	if err != nil {
		//log error
		log.Fatal(err)
		//stop
		return
	}

	/*
			hostname_current, err := get_hostname(raw_current_url)
			if err != nil {
				//log error
				log.Fatal(err)
				//stop
				return
			}


		//println("hostname_base: " + hostname_base)
		//println("hostname_current: " + hostname_current)
		//check if we are still on the some hostname
		if hostname_base != hostname_current {
			println("skipping " + hostname_current + "(not part of " + hostname_base + ")")
			//not same hostname: stop
			return
		}
	*/

	//get htmlBody from url
	htmlBody, err := getHTML(raw_current_url)
	if err != nil {
		//log error
		//log.Fatal(err)
		//stop
		return
	}
	//get urls
	url_list, err := getURLsFromHTML(htmlBody, raw_base_url)
	//if error
	if err != nil {
		log.Fatal(err)
		return
	}

	for i := 0; i < len(url_list); i++ {
		//get the url
		url_found := url_list[i]

		hostname_found, err := get_hostname(url_found)
		if err != nil {
			//log error
			log.Fatal(err)
			//stop
			return
		}
		//check if it is the same hostname
		if hostname_base != hostname_found {
			//println("skipping " + hostname_found + "(not part of " + hostname_base + ")")
			//not same hostname: stop
			return
		}

		//get the number of times this page was found (0 if not found)
		url_times_found := pages[url_found]
		//add the count
		pages[url_found] = url_times_found + 1
		//if found for the 1st time
		if url_times_found == 1 {
			//log
			//print("url found: " + url_found)
			//crawl this url
			crawlPage(raw_base_url, url_found, pages)
		}
	}
}
