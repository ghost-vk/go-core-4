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
