package main

import (
	"sort"
	"testing"
)

func Test_Strings(t *testing.T) {
	tt := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "3 simple strings",
			args: []string{"bca", "abc", "ab"},
			want: []string{"ab", "abc", "bca"},
		},
		{
			name: "cyrillic strings",
			args: []string{"абв", "вба", "бва"},
			want: []string{"абв", "бва", "вба"},
		},
	}
	for _, c := range tt {
		t.Run(c.name, func(t *testing.T) {
			sort.Strings(c.args)
			for i, want := range c.want {
				got := c.args[i]
				if got != want {
					t.Errorf("error sort string: got=%v, want=%v", got, want)
				}
			}
		})
	}
}
