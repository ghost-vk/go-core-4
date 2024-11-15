package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/go-core-4/gosearch/pkg/crawler"
	"github.com/go-core-4/gosearch/pkg/crawler/spider"
	"github.com/go-core-4/gosearch/pkg/index"
)

func main() {
	in := parseInput()

	spider := spider.New()
	pages1, err := scan(spider, "https://go.dev")
	if err != nil {
		fmt.Println("error go.dev scan page", err.Error())
	}

	idx := index.New()
	idx.Save(pages1)
	res := idx.Find(in.s)

	for _, doc := range res {
		fmt.Println(doc)
	}

	// pages2, err := scan(spider, "https://golang.org")
	// if err != nil {
	// 	fmt.Println("error scan golang.org page", err.Error())
	// }
	// pages2 := []crawler.Document{}

	// for _, p := range append(pages1, pages2...) {
	// 	if strings.Contains(strings.ToLower(p.Title), in.s) {
	// 		fmt.Println("ID=", p.ID, p.Title, "=>", p.URL)
	// 	}
	// }
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
