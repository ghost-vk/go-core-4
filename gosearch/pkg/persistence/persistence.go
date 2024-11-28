package persistence

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-core-4/gosearch/pkg/crawler"
)

type Persistence struct {
	file string
}

func New(file string) *Persistence {
	p := Persistence{}
	p.file = file
	return &p
}

func (p *Persistence) Documents() ([]crawler.Document, error) {
	f, err := os.OpenFile(p.file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		msg := fmt.Sprintf("error create or read cache file: %v", err.Error())
		return nil, errors.New(msg)
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		msg := fmt.Sprintf("error stat file: %v", err.Error())
		return nil, errors.New(msg)
	}

	docs, err := readDocuments(f, fstat.Size())
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (p *Persistence) Save(documents []crawler.Document) error {
	f, err := os.Create(p.file)
	if err != nil {
		msg := fmt.Sprintf("error create file: %v", err.Error())
		return errors.New(msg)
	}
	defer f.Close()

	data, err := serialize(documents)
	if err != nil {
		msg := fmt.Sprintf("error serialize data: %v", err.Error())
		return errors.New(msg)
	}

	err = write(f, data)
	if err != nil {
		msg := fmt.Sprintf("error write data to file: %v", err.Error())
		return errors.New(msg)
	}

	return nil
}

func write(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	if err != nil {
		msg := fmt.Sprintf("Error write to file: %v", err.Error())
		return errors.New(msg)
	}
	return nil
}

type documentMap struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func readDocuments(r io.Reader, size int64) ([]crawler.Document, error) {
	if size == 0 {
		return []crawler.Document{}, nil
	}
	data := make([]byte, size)
	for {
		_, err := r.Read(data)
		if err == io.EOF {
			break
		}
	}
	docs, err := deserialize(data)
	if err != nil {
		msg := fmt.Sprintf("Error deserialize: %v", err.Error())
		return nil, errors.New(msg)
	}
	return docs, nil
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
