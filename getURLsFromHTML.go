//part of https://www.boot.dev/lessons/98b7641d-2eae-4181-8a96-03f656a7315b

package main

import (
	"log"
	"strings"
	"golang.org/x/net/html" //https://pkg.go.dev/golang.org/x/net/html
	"golang.org/x/net/html/atom"
)

//get the urls from html body
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal(err)
	}
	url_list := []string{}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					//check if the URI can be normalizes
					url_found, err := normalizeURL_rawBase(a.Val, rawBaseURL)
					//if error
					if err != nil {
						//log error
						log.Fatal(err)
						//stop
						return []string{}, err
					}
					//add to url list
					url_list = append(url_list, url_found)
					//stop
					break
				}
			}
		}
	}
	//return the list
	return url_list, nil
}

/*
I'll try not to give too many hints: read the package docs linked above. Reading docs is vital practice, that said, here are a few hints:

    strings.NewReader(htmlBody) creates a io.Reader
    html.Parse(htmlReader) creates a tree of html.Nodes
    Use recursion to traverse the node tree and find the <a> tag "anchor" elements
    In HTML, "anchor" elements are links. e.g:
	<a href="https://www.boot.dev">Learn Backend Development</a>
*/
