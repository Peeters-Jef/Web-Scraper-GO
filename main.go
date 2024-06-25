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
    url, _, err := reader.ReadLine()
    if err != nil {
        fmt.Println("Error reading the input", err)
        return
    }

    queryURL := scraper.NormalizeURL(string(url))
    fmt.Println(queryURL)
}
