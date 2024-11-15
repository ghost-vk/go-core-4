package index

import (
	"strings"

	"github.com/go-core-4/gosearch/pkg/crawler"
)

// инвертированный индекс
type Index struct {
	items   map[string][]int
	storage []crawler.Document
}

// Хранение индексированных документов
// map[string]int
// map[string][]int
// map[string]map[int]bool

func New() *Index {
	idx := Index{}
	idx.items = make(map[string][]int)
	idx.storage = []crawler.Document{}
	return &idx
}

func (idx *Index) Save(docs []crawler.Document) {
	for docid, doc := range docs {
		idx.storage = append(idx.storage, doc)
		for _, word := range strings.Split(doc.Title, " ") {
			idx.items[word] = append(idx.items[word], docid)
		}
	}
}

func (idx *Index) Find(word string) []crawler.Document {
	items := idx.items[word]
	res := make([]crawler.Document, 0)
	for _, item := range items {
		res = append(res, idx.storage[item])
	}
	return res
}
