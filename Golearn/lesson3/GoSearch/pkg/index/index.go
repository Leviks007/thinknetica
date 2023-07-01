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

func NewIndex() *Index {
	return &Index{
		documents: []*crawler.Document{},
		indexMap:  make(map[string][]int),
	}
}

func (idx *Index) AddDocument(doc *crawler.Document) {

	words := strings.Fields(doc.Title)
	for _, word := range words {
		word = strings.ToLower(word)
		idx.indexMap[word] = append(idx.indexMap[word], doc.ID)
	}

	i := sort.Search(len(idx.documents), func(i int) bool {
		return idx.documents[i].ID >= doc.ID
	})

	idx.documents = append(idx.documents[:i], append([]*crawler.Document{doc}, idx.documents[i:]...)...)

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
