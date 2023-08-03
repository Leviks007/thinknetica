package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"lesson12/GoSearch/pkg/crawler"
	"lesson12/GoSearch/pkg/crawler/spider"
	indexDoc "lesson12/GoSearch/pkg/index"
	"lesson12/GoSearch/pkg/webapp"
)

var Index *indexDoc.Index

func main() {
	Index = indexDoc.New()
	urls := getURLs()
	documents := scanWebsites(urls)
	Index.AddDocuments(documents)
	sort.Sort(Index)

	StartWeb()
}

func StartWeb() {
	mapHttp := make(map[string]func(http.ResponseWriter, *http.Request))
	mapHttp["/"] = mainHandler
	mapHttp["/search"] = handlerSearch
	mapHttp["/index"] = handlerIndex
	mapHttp["/doc"] = handlerDoc

	webapp.OpenWeb(mapHttp)
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

func printMatchingURLs(documents []*crawler.Document) string {
	var builder strings.Builder

	for _, doc := range documents {
		fmt.Fprintf(&builder, "URL: %s Title: %s\n", doc.URL, doc.Title)
	}

	return builder.String()
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body><h2>Go Web App</h2>`))
	w.Write([]byte(`<ul>`))
	w.Write([]byte(`<li><a href="/search">/search</a> - Search for a keyword using the "s" parameter.</li>`))
	w.Write([]byte(`<li><a href="/doc">/doc</a> - Display a test representation of the documents.</li>`))
	w.Write([]byte(`<li><a href="/index">/index</a> - Display a test representation of the index.</li>`))
	w.Write([]byte(`</ul>`))
	w.Write([]byte(`</body></html>`))

}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	searchParam := r.URL.Query().Get("s")

	fmt.Fprintf(w, printMatchingURLs(getDocByWord(Index, searchParam)))
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Index.String())
}

func handlerDoc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Index.StringDoc())
}
