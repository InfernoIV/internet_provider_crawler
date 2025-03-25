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

	fmt.Printf("starting crawl of: %v\n", BASE_URL)
}

/*
https://gobyexample.com/command-line-arguments
go build -o crawler && ./crawler BASE_URL
go build -o crawler && ./crawler google.nl
*/
