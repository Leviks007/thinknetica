package index

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"lesson12/GoSearch/pkg/crawler"
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

func (idx *Index) StringDoc() string {
	var builder strings.Builder
	builder.WriteString("Documents:\n")
	for _, doc := range idx.documents {
		fmt.Fprintf(&builder, "URL: %s Title: %s\n", doc.URL, doc.Title)
	}

	return builder.String()
}

func (idx *Index) String() string {
	var builder strings.Builder

	builder.WriteString("Documents:\n")
	for _, doc := range idx.documents {
		fmt.Fprintf(&builder, "URL: %s Title: %s\n", doc.URL, doc.Title)
	}

	builder.WriteString("Index Map:\n")

	var sortedKeys []string
	for word := range idx.indexMap {
		sortedKeys = append(sortedKeys, word)
	}
	sort.Strings(sortedKeys)

	for _, word := range sortedKeys {
		ids := idx.indexMap[word]
		fmt.Fprintf(&builder, "Word: %s IDs: %v\n", word, ids)
	}

	return builder.String()
}

func (idx *Index) Swap(i, j int) {
	idx.documents[i], idx.documents[j] = idx.documents[j], idx.documents[i]
}

func New() *Index {
	return &Index{
		documents: make([]*crawler.Document, 0),
		indexMap:  make(map[string][]int),
	}
}

func (idx *Index) AddDocuments(docs []crawler.Document) {
	for i, doc := range docs {
		if containsElementByURL(idx.documents, doc.URL) {
			continue
		}
		idx.documents = append(idx.documents, &docs[i])
		words := strings.Fields(doc.Title)
		for _, word := range words {
			word = strings.ToLower(word)
			if !findElement(idx.indexMap[word], doc.ID) {
				idx.indexMap[word] = append(idx.indexMap[word], doc.ID)
			}
		}
	}
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

func stringToID(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

//пока так дальше если их хранить в базе то будет другой способ этот не оптимальный
func containsElementByURL(arr []*crawler.Document, url string) bool {
	for _, doc := range arr {
		if doc.URL == url {
			return true
		}
	}
	return false
}
