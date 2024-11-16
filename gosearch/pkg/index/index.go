package index

import (
	"strings"

	"github.com/go-core-4/gosearch/pkg/crawler"
)

// инвертированный индекс
type Index struct {
	items map[string][]int
}

func New() *Index {
	idx := Index{}
	idx.items = make(map[string][]int)
	return &idx
}

func (idx *Index) Save(docs []crawler.Document) {
	for _, doc := range docs {
		for _, word := range strings.Split(doc.Title, " ") {
			idx.items[word] = append(idx.items[word], doc.ID)
		}
	}
}

func (idx *Index) Find(word string) []int {
	return idx.items[word]
}
