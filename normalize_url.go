package main

import "net/url"

func normalizeURL(url_input string) (string, error) {
	//parse the URL
	parsedUrl, err := url.Parse(url_input)
	//if error
	if err != nil {
		//stop
		return "", err
	}
	//create the normalized URL
	normalized_URL := parsedUrl.Hostname() + parsedUrl.EscapedPath()
	//print("url_input: " + url_input)
	//print("normalized_URL: " + normalized_URL)
	//return the normalized url
	return normalized_URL, nil
}
