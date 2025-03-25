package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	//get information
	resp, err := http.Get(rawURL)

	//if error
	if err != nil {
		//log error
		log.Fatal(err)
		//return error
		return "", err
	}

	//get status code
	status_code := resp.StatusCode
	//if error code
	if status_code >= 400 {
		//log error
		log.Fatal("http " + resp.Status)
		//return the error code
		return "", errors.New("ERROR: http " + resp.Status)
	}

	//get content type
	content_type := resp.Header.Get("Content-Type")
	//check if the content type is correct
	if !strings.Contains(content_type, "text/html") {
		//log error
		//log.Fatal("Content-Type: " + content_type)
		//return the error code
		return "", errors.New("ERROR: Content-Type: " + content_type)
	}

	//read the data
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%s", b)
	//stub
	return string(body), nil
}
