package main

import (
	"fmt"
	"net/url"
	"os"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	fmt.Printf("Now crawling %s...\n", rawCurrentURL)
	if b, _ := shareSameDomain(rawBaseURL, rawCurrentURL); !b {
		return
	}
	normCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not normalize url %s: %v\n", rawCurrentURL, err)
	}
	pages[normCurrentUrl]++
	if pages[normCurrentUrl] > 3 {
		fmt.Printf("Too many access request to %s\n", normCurrentUrl)
		fmt.Println(pages)
		os.Exit(1)
	}
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not parse %s: %v\n", htmlBody, err)
	}
	urls, err := getURLsFromHTML(htmlBody, rawCurrentURL)
	if err != nil {
		fmt.Printf("could not get urls from html: %v\n", err)
	}
	for _, url := range urls {
		normUrl, err := normalizeURL(url)
		if err != nil {
			fmt.Printf("could not normalize url %s: %v\n", normUrl, err)
		}
		if pages[normUrl] == 0 {
			crawlPage(rawBaseURL, url, pages)
		}
	}
}

func shareSameDomain(url1, url2 string) (bool, error) {
	firstUrl, err := url.Parse(url1)
	if err != nil {
		return false, fmt.Errorf("coul not parse url %s: %v", url1, err)
	}
	secondUrl, err := url.Parse(url2)
	if err != nil {
		return false, fmt.Errorf("coul not parse url %s: %v", url2, err)
	}
	return firstUrl.Host == secondUrl.Host, nil
}
