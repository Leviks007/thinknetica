package main

import (
	"flag"
	"fmt"
	"lesson5/GoSearch/pkg/crawler"
	"lesson5/GoSearch/pkg/crawler/spider"
	indexDoc "lesson5/GoSearch/pkg/index"
	"log"
	"os"
	"sort"
)

func main() {
	searchKeyword := flag.String("s", "", "Параметр поиска")
	flag.Parse()

	f, err := os.OpenFile("index.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	FF := true //file found
	indexFF, err := indexDoc.GetIndexFromFile(f)
	if err != nil {
		log.Printf("Ошибка чтения файла")
		FF = false
	}
	var index *indexDoc.Index
	if FF && len(indexFF.Documents) > 0 {
		index = indexFF
	} else {
		urls := getURLs()
		documents := scanWebsites(urls)
		index = indexDoc.New()
		index.AddDocuments(documents)
		sort.Sort(index)
		err = index.WriteIndexToJson(f)
		if err != nil {
			fmt.Println("ошибка записи файла:", err)
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
