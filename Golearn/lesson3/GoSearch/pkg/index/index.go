package index

import (
	"lesson3/GoSearch/pkg/crawler"
	"sort"
	"strings"
)

type Index struct {
	documents []*crawler.Document
	indexMap  map[string][]int
}

func (idx *Index) Len() int {
	return len(idx.documents)
}

func (idx *Index) Less(i, j int) bool {
	return idx.documents[i].ID < idx.documents[j].ID
}

func (idx *Index) Swap(i, j int) {
	idx.documents[i], idx.documents[j] = idx.documents[j], idx.documents[i]
}

func New() *Index {
	return &Index{
		documents: []*crawler.Document{},
		indexMap:  make(map[string][]int),
	}
}

func (idx *Index) AddDocument(doc *crawler.Document) {
	words := strings.Fields(doc.Title)
	for _, word := range words {
		word = strings.ToLower(word)
		if !findElement(idx.indexMap[word], doc.ID) {
			idx.indexMap[word] = append(idx.indexMap[word], doc.ID)
		}
	}
	idx.documents = append(idx.documents, doc)
}

func (idx *Index) Search(word string) []int {
	word = strings.ToLower(word)
	return idx.indexMap[word]
}

func (idx *Index) GetDocsByID(ids []int) []*crawler.Document {
	var docs []*crawler.Document
	for _, id := range ids {
		i := sort.Search(len(idx.documents), func(i int) bool {
			return idx.documents[i].ID >= id
		})
		docs = append(docs, idx.documents[i])
	}
	return docs
}

func findElement(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
