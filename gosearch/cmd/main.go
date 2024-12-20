package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/go-core-4/gosearch/pkg/crawler"
	"github.com/go-core-4/gosearch/pkg/crawler/spider"
	"github.com/go-core-4/gosearch/pkg/index"
	"github.com/go-core-4/gosearch/pkg/persistence"
)

type MemStorage struct {
	documents []crawler.Document
}

func main() {
	in := parseInput()
	// for debug
	// _ = parseInput()

	storage := MemStorage{
		documents: make([]crawler.Document, 0),
	}

	persFile := "./cache.json"
	pers := persistence.New(persFile)
	documents, err := pers.Documents()
	if err != nil {
		log.Printf("Error read documents from file %v: %v", persFile, err)
	}

	if len(documents) == 0 {
		spider := spider.New()
		godevDocs, err := scan(spider, "https://go.dev")
		if err != nil {
			fmt.Println("error go.dev scan page", err.Error())
		}
		golangDocs, err := scan(spider, "https://golang.org")
		if err != nil {
			fmt.Println("error scan golang.org page", err.Error())
		}
		storage.documents = append(storage.documents, append(godevDocs, golangDocs...)...)
		pers.Save(storage.documents)
	} else {
		storage.documents = documents
	}

	sort.Sort(ById(storage.documents))

	idx := index.New()
	idx.Save(storage.documents)

	docIds := idx.Find(in.s)
	// for debug
	// docIds := idx.Find("package")

	for _, id := range docIds {
		doc, _ := binarySearch(id, storage.documents)
		fmt.Println(doc)
	}
}

func scan(searchService *spider.Service, url string) ([]crawler.Document, error) {
	pages, err := searchService.Scan(url, 2)

	if err != nil {
		fmt.Printf("error %s scan page: %v\n", url, err.Error())
	}

	return pages, nil
}

func parseInput() Input {
	search := flag.String("s", "", "search word")
	flag.Parse()

	return Input{
		s: strings.ToLower(*search),
	}
}

type Input struct {
	s string
}

type ById []crawler.Document

func (items ById) Len() int {
	return len(items)
}

func (items ById) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items ById) Less(i, j int) bool {
	return items[i].ID < items[j].ID
}

func binarySearch(targetId int, documents []crawler.Document) (crawler.Document, error) {
	from := 0
	to := len(documents) - 1
	for to >= from {
		mid := (from + to) / 2
		doc := documents[mid]
		if doc.ID == targetId {
			return doc, nil
		}
		if targetId > doc.ID {
			from = mid + 1
		} else {
			to = mid - 1
		}
	}

	return crawler.Document{}, errors.New("document not found")
}
