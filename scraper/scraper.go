package scraper

import (
	//"net/http"
	"net/url"
	"strings"
)

func NormalizeURL(rawUrl string) string {
    if !strings.Contains(rawUrl, "://") {
        rawUrl = "https://" + rawUrl
    }
    
    parsedURL, err := url.Parse(rawUrl)
    if err != nil {
        return "no url provided"
    }

    if parsedURL.Scheme != "https" {
        parsedURL.Scheme = "https"
    }

    parsedURL.Host = parsedURL.Hostname()
    parsedURL.Path = "/"

    return parsedURL.String()
}

