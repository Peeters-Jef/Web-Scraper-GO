package main

import (
	"bufio"
	"fmt"
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

    visitedLinks := make(map[string]bool)
    scraper.Crawl(string(baseURL), visitedLinks)
}
