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

func Crawl(urlString, baseDomain string, visited map[string]bool) {
    if visited[urlString] {
        return
    }

    resp, err := http.Get(urlString)
    if err != nil {
        fmt.Println("Get Error:", err)
        return
    }

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Status Error:", resp.StatusCode)
        return
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    }

    visited[urlString] = true
    fmt.Println("Visiting:", urlString)

    baseURL := urlString
    links := extractLinks(doc)

    for _, link := range links {
        resolvedLink := resolveURL(link, baseURL)
        normalizedURL := NormalizeURL(resolvedLink)

        parsedResolvedLink, err := url.Parse(resolvedLink)
        if err != nil || parsedResolvedLink.Host != baseDomain {
            continue
        }

        if !visited[normalizedURL] {
            Crawl(resolvedLink, baseDomain, visited)
        }
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

func resolveURL(href, baseURL string) string {
    u, err := url.Parse(href)
    if err != nil {
        return ""
    }
    baseU, err := url.Parse(baseURL)
    if err != nil {
        return ""
    }

    return baseU.ResolveReference(u).String()
}
