package main

import (
	"fmt"
	"webscraper/scraper"
)

func tess(url string) {
    
    res := scraper.NormalizeURL(url)
    fmt.Println(res)
}

func main() {
    url := "wagslane.dev/"
    tess(url)
}
