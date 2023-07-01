package main

import (
	"flag"
	"fmt"
	"lesson2/pkg/crawler"
	"lesson2/pkg/crawler/spider"
	"log"
	"strings"
)

func main() {
	a := 2
	_ = a
	searchKeyword := flag.String("s", "", "Параметр поиска")
	flag.Parse()
	urls := getURLs()
	documents := scanWebsites(urls)

	if *searchKeyword != "" {
		printMatchingURLs(documents, *searchKeyword)
	}
}

func getURLs() []string {
	return []string{"https://go.dev", "https://golang.org"}
}

func scanWebsites(urls []string) []crawler.Document {
	var documents []crawler.Document
	s := spider.New()

	for _, url := range urls {
		doc, err := s.Scan(url, 2)
		if err != nil {
			log.Printf("Ошибка при сканировании сайта %s: %v", url, err)
			continue
		}
		documents = append(documents, doc...)
	}

	return documents
}

func printMatchingURLs(documents []crawler.Document, keyword string) {
	for _, doc := range documents {
		if strings.Contains(doc.Title, keyword) {
			fmt.Println("URL:", doc.URL)
		}
	}
}
