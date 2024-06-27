package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func NormalizeURL(rawUrl string) string {
    parsedURL, err := url.Parse(rawUrl)
    if err != nil {
        return "no url provided"
    }

    // remove the Scheme
    normalizedURL := parsedURL.Host + parsedURL.Path

    normalizedURL = strings.ToLower(normalizedURL)

    if strings.HasPrefix(normalizedURL, "//") {
        return normalizedURL[2:]
    }
    
    if strings.HasSuffix(normalizedURL, "/") {
        return normalizedURL[:len(normalizedURL)-1]
    }

    return normalizedURL
}

func Crawl(url string, visited map[string]bool) {
    if visited[url] {
        return
    }

    visited[url] = true
    fmt.Println("Visiting:", url)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error:", resp.StatusCode)
        return
    }
}

func extractLinks(n *html.Node) []string {
    var links []string

    if n.Type == html.ElementNode && n.Data == "a" {
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                links = append(links, attr.Val)
            }
        }
    }
    for child := n.FirstChild; child != nil; child = child.NextSibling {
        links = append(links, extractLinks(child)...)
    } 
    return links
}
