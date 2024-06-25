package scraper

import (
	"net/url"
	"strings"
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
