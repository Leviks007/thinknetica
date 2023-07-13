package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"lesson5/GoSearch/pkg/crawler"
	"lesson5/GoSearch/pkg/crawler/spider"
	indexDoc "lesson5/GoSearch/pkg/index"
)

func main() {
	namef := "index.json"
	searchKeyword := flag.String("s", "", "Параметр поиска")
	flag.Parse()

	var index *indexDoc.Index
	if doesFileExist(namef) {
		f, err := os.Open(namef)
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		indexFF, err := indexDoc.GetIndexFromFile(f)
		if err != nil {
			log.Println(err)
			return
		}
		err = json.Unmarshal(indexFF, &index)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if index.IsEmpty() {
		urls := getURLs()
		documents := scanWebsites(urls)
		index = indexDoc.New()
		index.AddDocuments(documents)
		sort.Sort(index)

		f, err := os.Create(namef)
		if err != nil {
			log.Println("Ошибка при создании файла:", err)
			return
		}
		defer f.Close()

		err = index.WriteIndexToJson(f)
		if err != nil {
			log.Println("ошибка записи файла:", err)
			return
		}
	}

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
		doc, err := s.Scan(url, 1)
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

func doesFileExist(path string) bool {
	found := false
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			log.Println(err)
		}
	} else {
		found = true
	}
	return found
}
