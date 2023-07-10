package index

import (
	"encoding/json"
	"io"
	"lesson5/GoSearch/pkg/crawler"
	"sort"
	"strings"
)

type Index struct {
	Documents []*crawler.Document `json:"documents"`
	IndexMap  map[string][]int    `json:"indexMap"`
}

func (idx *Index) Len() int {
	return len(idx.Documents)
}

func (idx *Index) Less(i, j int) bool {
	return idx.Documents[i].ID < idx.Documents[j].ID
}

func (idx *Index) Swap(i, j int) {
	idx.Documents[i], idx.Documents[j] = idx.Documents[j], idx.Documents[i]
}

func New() *Index {
	return &Index{
		Documents: []*crawler.Document{},
		IndexMap:  make(map[string][]int),
	}
}

func (idx *Index) AddDocuments(docs []crawler.Document) {
	idx.Documents = make([]*crawler.Document, len(docs))
	for i, doc := range docs {
		idx.Documents[i] = &docs[i]
		words := strings.Fields(doc.Title)
		for _, word := range words {
			word = strings.ToLower(word)
			if !findElement(idx.IndexMap[word], doc.ID) {
				idx.IndexMap[word] = append(idx.IndexMap[word], doc.ID)
			}
		}
	}
}

func (idx *Index) Search(word string) []int {
	word = strings.ToLower(word)
	return idx.IndexMap[word]
}

func (idx *Index) GetDocsByID(ids []int) []*crawler.Document {
	var docs []*crawler.Document
	for _, id := range ids {
		i := sort.Search(len(idx.Documents), func(i int) bool {
			return idx.Documents[i].ID >= id
		})
		docs = append(docs, idx.Documents[i])
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

func GetIndexFromFile(r io.Reader) ([]byte, error) {
	fileData, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return fileData, err
}

func (idx *Index) WriteIndexToJson(w io.Writer) error {
	jsonData, err := json.Marshal(idx)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonData)
	return err
}

func (idx *Index) IsEmpty() bool {
	return idx == nil || (len(idx.Documents) == 0 && len(idx.IndexMap) == 0)
}
