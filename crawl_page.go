package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	// Do not crawl outside websites
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not normalize url %s: %v\n", currentURL, err)
	}
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		// If not first visit, adding page visit to config is enough
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not get html %s: %v\n", htmlBody, err)
	}
	nextUrls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("could not get urls from html: %v\n", err)
	}
	for _, nextUrl := range nextUrls {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextUrl)
	}
}
