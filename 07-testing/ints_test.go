package main

import (
	"sort"
	"testing"
)

func Test_Ints(t *testing.T) {
	wantInts := []int{1, 2, 3}
	gotInts := []int{3, 2, 1}

	if len(wantInts) != len(gotInts) {
		t.Error("error sort ints: bad test data, got and want slice length should be equal")
		t.FailNow()
	}

	sort.Ints(gotInts)
	for i, want := range wantInts {
		got := gotInts[i]
		if got != want {
			t.Errorf("error sort ints: got %v, want %v", got, wantInts)
		}
	}
}
