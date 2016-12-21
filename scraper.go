package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getPages(urls []string) (pageData map[string]string) {
	pageData = make(map[string]string)
	for _, url := range urls {
		response, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(response.Body)
		pageData[url] = string(bytes)
		response.Body.Close()
	}
	return
}

func main() {
	urls := []string{"https://en.wikipedia.org/wiki/Page", "https://en.wikipedia.org/wiki/Page"}
	pages := getPages(urls)
	links := parseLinks(pages)
	for url, links := range links {
		fmt.Println(url, links)
	}
}

func parseLinks(pageData map[string]string) (linkSlice map[string][]string) {
	linkSlice = make(map[string][]string)
	for url, body := range pageData {
		linkSlice[url] = strings.Split(body, "<a href=")[1:]
	}
	for url, links := range linkSlice {
		linkUrls := []string{}
		for _, link := range links {
			linkUrls = append(linkUrls, strings.Split(link, "\"")[1]) // idx 0 or 1?
		}
		linkSlice[url] = linkUrls
	}
	return
}
