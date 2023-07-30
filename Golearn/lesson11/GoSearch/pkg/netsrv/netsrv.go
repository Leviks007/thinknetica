// В файле netsrv.go
package netsrv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sort"
	"time"

	"lesson11/GoSearch/pkg/crawler"
	"lesson11/GoSearch/pkg/crawler/spider"
	indexDoc "lesson11/GoSearch/pkg/index"
)

func handler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("Conn closed")

	conn.SetDeadline(time.Now().Add(time.Minute * 10))

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		urls := getURLs()
		documents := scanWebsites(urls)
		index := indexDoc.New()
		index.AddDocuments(documents)
		sort.Sort(index)

		answerM := printMatchingURLs(getDocByWord(index, string(msg)))
		answerM = append(answerM, '\n')
		_, err = conn.Write(answerM)
		if err != nil {
			return
		}
		conn.SetDeadline(time.Now().Add(time.Minute * 10))
	}
}

func StartServer() {
	// регистрация сетевой службы
	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on port 8000...")
	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Conn established")

		go handler(conn)
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

func printMatchingURLs(documents []*crawler.Document) []byte {
	var results []URLWithTitle

	for _, doc := range documents {
		result := URLWithTitle{
			URL:   doc.URL,
			Title: doc.Title,
		}
		results = append(results, result)
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil
	}

	return jsonData
}

type URLWithTitle struct {
	URL   string `json:"URL"`
	Title string `json:"Title"`
}
