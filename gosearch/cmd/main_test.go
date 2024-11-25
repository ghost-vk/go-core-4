package main

import (
	"reflect"
	"testing"

	"github.com/go-core-4/gosearch/pkg/crawler"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		docId     int
		documents []crawler.Document
	}
	docs := make([]crawler.Document, 0)
	doc := crawler.Document{
		ID: 1,
	}
	doc2 := crawler.Document{
		ID: 2,
	}
	doc3 := crawler.Document{
		ID: 3,
	}
	docs = append(docs, doc, doc2, doc3)
	tests := []struct {
		name string
		args args
		want crawler.Document
	}{
		{
			name: "success primitive case",
			args: args{
				docId:     1,
				documents: docs,
			},
			want: doc,
		},
		{
			name: "one item case",
			args: args{
				docId:     1,
				documents: docs[1:2],
			},
			want: doc2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := binarySearch(tt.args.docId, tt.args.documents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serialize(t *testing.T) {
	type args struct {
		docs []crawler.Document
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Single valid document",
			args: args{
				docs: []crawler.Document{
					{
						ID:    1,
						URL:   "http://example.com",
						Title: "Example",
						Body:  "This is an example document.",
					},
				},
			},
			want:    []byte("\n{\"ID\":1,\"URL\":\"http://example.com\",\"Title\":\"Example\",\"Body\":\"This is an example document.\"}"),
			wantErr: false,
		},
		{
			name: "Multiple valid documents",
			args: args{
				docs: []crawler.Document{
					{
						ID:    1,
						URL:   "http://example.com",
						Title: "Example",
						Body:  "This is an example document.",
					},
					{
						ID:    2,
						URL:   "http://test.com",
						Title: "Test",
						Body:  "This is another document.",
					},
				},
			},
			want:    []byte("\n{\"ID\":1,\"URL\":\"http://example.com\",\"Title\":\"Example\",\"Body\":\"This is an example document.\"}\n{\"ID\":2,\"URL\":\"http://test.com\",\"Title\":\"Test\",\"Body\":\"This is another document.\"}"),
			wantErr: false,
		},
		{
			name: "Empty input",
			args: args{
				docs: []crawler.Document{},
			},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "Document with special characters",
			args: args{
				docs: []crawler.Document{
					{
						ID:    1,
						URL:   "http://example.com",
						Title: "Title with \"quotes\" and \\slashes\\",
						Body:  "Body with newline\nand tab\tcharacters.",
					},
				},
			},
			want:    []byte("\n{\"ID\":1,\"URL\":\"http://example.com\",\"Title\":\"Title with \\\"quotes\\\" and \\\\slashes\\\\\",\"Body\":\"Body with newline\\nand tab\\tcharacters.\"}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serialize(tt.args.docs)
			if (err != nil) != tt.wantErr {
				t.Errorf("serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
