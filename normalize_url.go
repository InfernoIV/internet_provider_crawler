package main

import (
	"log"
	"net/url"
)

// function that checks and fixes the url
func normalize_url(url_input string) (string, error) {
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
	//normalized_URL := parsedUrl.Scheme + "://" + parsedUrl.Hostname() + "/" + parsedUrl.EscapedPath()
	//return the normalized url
	return parsedUrl.String(), nil //normalized_URL, nil
}

// function that parses urls (and also only paths)
func normalize_url_raw_base(url_input string, rawBaseURL string) (string, error) {
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

	//create normalized url
	normalized_url := url.URL{}

	//set hostname
	normalized_url.Host = parsed_url_input.Hostname()
	//check if we need to replace
	if (normalized_url.Hostname() == "") {
		//use the hostname of the base url
		normalized_url.Host = parsed_base_url.Hostname()
	}

	//set scheme
	normalized_url.Scheme = parsed_url_input.Scheme
	//check if scheme exists
	if (normalized_url.Scheme == "") {
		//use the scheme of the base url
		normalized_url.Scheme = parsed_base_url.Scheme
	}

	//set path
	normalized_url.Path = parsed_url_input.Path

	//return the created url
	return normalized_url.String(), nil

	/*
	type URL ¶
	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
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
}



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
