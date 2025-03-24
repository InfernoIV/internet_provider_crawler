//part of https://www.boot.dev/lessons/98b7641d-2eae-4181-8a96-03f656a7315b

package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	r := strings.NewReader(htmlBody)
	z := html.NewTokenizer(r)

	/* TODO
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			// ...
			return ...
		}
		// Process the current token.
	}
	*/
	print(z)
	print(rawBaseURL)
	//nodes := html.Parse(reader)

	return nil, nil
}

/*
I'll try not to give too many hints: read the package docs linked above. Reading docs is vital practice, that said, here are a few hints:

    strings.NewReader(htmlBody) creates a io.Reader
    html.Parse(htmlReader) creates a tree of html.Nodes
    Use recursion to traverse the node tree and find the <a> tag "anchor" elements
    In HTML, "anchor" elements are links. e.g:
	<a href="https://www.boot.dev">Learn Backend Development</a>
*/
