package main

import (
	"flag"
	"fmt"
	"lesson3/GoSearch/pkg/crawler"
	"lesson3/GoSearch/pkg/crawler/spider"
	indexDoc "lesson3/GoSearch/pkg/index"
	"log"
	"sort"
)

func main() {

	searchKeyword := flag.String("s", "", "Параметр поиска")
	flag.Parse()
	urls := getURLs()
	documents := scanWebsites(urls)
	index := indexDoc.New()
	index.AddDocuments(documents)
	sort.Sort(index)

	if *searchKeyword != "" {
		printMatchingURLs(getDocByWord(index, *searchKeyword))
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

func getDocByWord(idx *indexDoc.Index, searchKeyword string) []*crawler.Document {
	IDs := idx.Search(searchKeyword)
	return idx.GetDocsByID(IDs)
}

func printMatchingURLs(documents []*crawler.Document) {
	for _, doc := range documents {
		fmt.Println(
			"URL:",
			doc.URL,
			"Title:",
			doc.Title)
	}
}
