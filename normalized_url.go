package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(s string) (string, error) {
	parsedUrl, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("could not parse url: %v", err)
	}
	fullUrl := parsedUrl.Host + parsedUrl.Path
	fullUrl = strings.ToLower(fullUrl)
	fullUrl = strings.TrimSuffix(fullUrl, "/")
	return fullUrl, nil
}
