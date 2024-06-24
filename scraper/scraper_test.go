package scraper

import (
    "testing"
)

func TestNormalizeURLStripProtocol(t *testing.T) {
    input := "https://blog.boot.dev"
    actual := NormalizeURL(input)
    want := "blog.boot.dev"

    if actual != want {
        t.Fatalf(`%v is not the same as %v, the protocol was not removed`, actual, want)
    }
}

func TestNormalizeURLStripSlash(t *testing.T) {
    input := "blog.boot.dev/"
    actual := NormalizeURL(input)
    want := "blog.boot.dev"

    if actual != want {
        t.Fatalf(`%v is not the same as %v, the slash was not removed`, actual, want)
    }
}

func TestNormalizeURLCapitals(t *testing.T) {
    input := "blog.boot.dev/"
    actual := NormalizeURL(input)
    want := "blog.boot.dev"

    if actual != want {
        t.Fatalf(`%v is not the same as %v, the capital letters were not removed`, actual, want)
    }
}

