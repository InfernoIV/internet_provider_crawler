package main

import (
	"log"
	"net/url"
)

func normalizeURL(url_input string) (string, error) {
	//parse the URL
	parsedUrl, err := url.Parse(url_input)
	//debug
	//print(fmt.Sprintf("parsedUrl is %v", parsedUrl))
	//if error
	if err != nil {
		//log error
		log.Fatal(err)
		//stop
		return "", err
	}
	
	//create the normalized URL
	normalized_URL := parsedUrl.Scheme + "://" + parsedUrl.Hostname() + parsedUrl.EscapedPath() 
	//return the normalized url
	return normalized_URL, nil
}

func normalizeURL_rawBase(url_input string, rawBaseURL string) (string, error) {
	//parse the base URL
	parsed_base_url, err := url.Parse(rawBaseURL)
	//if error
	if err != nil {
		//log error
		log.Fatal(err)
		//stop
		return "", err
	}

	//parse the input URL
	parsed_url_input, err := url.Parse(url_input)
	//if error
	if err != nil {
		//log error
		log.Fatal(err)
		//stop
		return "", err
	}

	//get scheme
	scheme := parsed_url_input.Scheme
	//check scheme
	if scheme == "" {
		scheme = parsed_base_url.Scheme
	}
	//get hostname
	hostname := parsed_url_input.Hostname()
	//check if it is relative
	if hostname == "" {
		hostname = parsed_base_url.Hostname()
	}
	return_url := scheme + "://" + hostname + parsed_url_input.EscapedPath()
	//fmt.Println("Got '", url_input, "', created '", return_url, "'")
	return return_url, nil
}

/*
 type URL ¶

type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port (see Hostname and Port methods)
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	OmitHost    bool      // do not emit empty host (authority)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}


*/

func get_hostname(url_input string) (string, error) {
	//parse the URL
	parsedUrl, err := url.Parse(url_input)
	//if error
	if err != nil {
		//log error
		log.Fatal(err)
		//stop
		return "", err
	}
	//only return the hostname
	return parsedUrl.Hostname(), nil
}
