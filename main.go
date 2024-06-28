package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"webscraper/scraper"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Please enter the URL: ")
    baseURL, _, err := reader.ReadLine()
    if err != nil {
        fmt.Println("Error reading the input", err)
        return
    }

    parsedURL, err := url.Parse(string(baseURL))
    baseDomain := parsedURL.Host

    visitedLinks := make(map[string]bool)
    scraper.Crawl(string(baseURL), baseDomain, visitedLinks)
}
