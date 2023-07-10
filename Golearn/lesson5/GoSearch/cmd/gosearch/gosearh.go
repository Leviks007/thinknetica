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
	searchKeyword := flag.String("s", "", "Параметр поиска")
	flag.Parse()

	f, err := os.OpenFile("index.json", os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Файл не найден, будет создан новый!")
	}
	defer f.Close()

	var index *indexDoc.Index
	if f != nil {
		indexFF, err := indexDoc.GetIndexFromFile(f)
		if err != nil {
			log.Printf("Ошибка чтения файла")
		}
		err = json.Unmarshal(indexFF, &index)
		if err != nil {
			log.Printf("Ошибка преобразования файла")
		}
	}
	if index.IsEmpty() {
		urls := getURLs()
		documents := scanWebsites(urls)
		index = indexDoc.New()
		index.AddDocuments(documents)
		sort.Sort(index)

		f, err := os.Create("index.json")
		if err != nil {
			log.Fatal("Ошибка при создании файла:", err)
		}
		defer f.Close()
		err = index.WriteIndexToJson(f)
		if err != nil {
			log.Fatal("ошибка записи файла:", err)
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
