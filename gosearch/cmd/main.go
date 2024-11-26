package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/go-core-4/gosearch/pkg/crawler"
	"github.com/go-core-4/gosearch/pkg/crawler/spider"
	"github.com/go-core-4/gosearch/pkg/index"
)

type Storage struct {
	documents []crawler.Document
}

func main() {
	in := parseInput()

	storage := Storage{
		documents: make([]crawler.Document, 0),
	}

	f, err := os.OpenFile("./cache.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Error create or read cache file: %v", err.Error())
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Printf("Error stat file: %v", err.Error())
		return
	}
	restore(f, &storage, fi.Size())

	if len(storage.documents) == 0 {
		spider := spider.New()
		godevDocs, err := scan(spider, "https://go.dev")
		if err != nil {
			fmt.Println("error go.dev scan page", err.Error())
		}

		// For debug
		// godevDocs := []crawler.Document{
		// 	{ID: 25375595, URL: "https://go.dev/learn#featured-books", Title: "Get Started - The Go Programming Language", Body: ""},
		// 	{ID: 81644140, URL: "https://go.dev/pkg", Title: "Standard library - Go Packages", Body: ""},
		// 	{ID: 77574494, URL: "https://go.dev/conduct", Title: "Go Community Code of Conduct - The Go Programming Language", Body: ""},
		// }

		golangDocs, err := scan(spider, "https://golang.org")
		if err != nil {
			fmt.Println("error scan golang.org page", err.Error())
		}

		f, err := os.Create("./cache.txt")
		if err != nil {
			log.Printf("Error create file: %v", err.Error())
			return
		}
		defer f.Close()

		storage.documents = append(storage.documents, append(godevDocs, golangDocs...)...)

		serialized, err := serialize(storage.documents)
		if err != nil {
			log.Printf("Error serialize results: %v", err.Error())
			return
		}

		cache(f, serialized)
	}

	sort.Sort(ById(storage.documents))

	idx := index.New()
	idx.Save(storage.documents)

	docIds := idx.Find(in.s)
	// For debug
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

func cache(w io.Writer, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Printf("Error write to file: %v", err.Error())
		return
	}
}

func restore(r io.Reader, s *Storage, size int64) {
	data := make([]byte, size)
	for {
		_, err := r.Read(data)
		if err == io.EOF {
			break
		}
	}
	docs, err := deserialize(data)
	if err != nil {
		log.Printf("Error deserialize: %v", err.Error())
		return
	}
	s.documents = docs
}

type documentMap struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func serialize(docs []crawler.Document) ([]byte, error) {
	buf := []documentMap{}
	for _, doc := range docs {
		buf = append(buf, documentMap{
			ID:    doc.ID,
			URL:   doc.URL,
			Title: doc.Title,
			Body:  doc.Body,
		})
	}
	result, err := json.Marshal(buf)
	if err != nil {
		log.Printf("Error serialize: %v", err.Error())
		return nil, errors.New("error serialize document")
	}
	return result, nil
}

func deserialize(raw []byte) ([]crawler.Document, error) {
	var parsed []documentMap
	err := json.Unmarshal(raw, &parsed)
	if err != nil {
		log.Printf("Error deserialize: %v", err.Error())
		return nil, errors.New("error deserialize document")
	}
	buf := []crawler.Document{}
	for _, p := range parsed {
		buf = append(buf, crawler.Document{
			ID:    p.ID,
			URL:   p.URL,
			Title: p.Title,
			Body:  p.Body,
		})
	}
	return buf, nil
}
