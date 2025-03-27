package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	//"fmt"
)

//Function that retrieves html from an url
func get_html(rawURL string) (string, error) {
	//get information
	resp, err := http.Get(rawURL)

	//if error
	if err != nil {
		//return error
		return "", err
	}

	//get status code
	status_code := resp.StatusCode
	//if error code
	if status_code >= 400 {
		//return the error code
		return "", errors.New("ERROR: http " + resp.Status)
	}

	//get content type
	content_type := resp.Header.Get("Content-Type")
	//check if the content type is correct
	if !strings.Contains(content_type, "text/html") {
		//return the error code
		return "", errors.New("ERROR: Content-Type: " + content_type)
	}

	//close reader when done
	defer resp.Body.Close()
	//start reading
	body, err := io.ReadAll(resp.Body)
	//if error
	if err != nil {
		//return error
		return "", err
	}

	//return the body
	return string(body), nil
}
