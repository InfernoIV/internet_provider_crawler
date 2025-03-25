package main

import (
	"fmt"
	"os"
)

func main() {
	//check if we have enough arguments (should be more than 1)
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	//check if we only have 1 argument
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	//get the url
	BASE_URL := os.Args[1]

	//start message
	fmt.Printf("starting crawl of: %v\n", BASE_URL)
	//get data
	data, err := getHTML(BASE_URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	print(data)
}

/*
https://gobyexample.com/command-line-arguments
go build -o crawler && ./crawler BASE_URL
go build -o crawler && ./crawler google.nl
go build -o crawler && ./crawler https://wagslane.dev
*/
