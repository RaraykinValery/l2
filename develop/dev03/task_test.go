package main

import (
	"reflect"
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		lines   []string
		numeric bool
		column  uint
		reverse bool
		want    []string
	}{
		{
			lines: []string{"b", "c", "a", "d"},
			want:  []string{"a", "b", "c", "d"},
		},
		{
			lines:   []string{"2", "1", "4", "3"},
			numeric: true,
			want:    []string{"1", "2", "3", "4"},
		},
		{
			lines:   []string{"b", "c", "a", "d"},
			reverse: true,
			want:    []string{"d", "c", "b", "a"},
		},
	}

	for _, tt := range tests {
		SortLines(tt.lines, tt.numeric, tt.column, tt.reverse)
		if !reflect.DeepEqual(tt.lines, tt.want) {
			t.Errorf("Strings %#v are not equal to %#v", tt.lines, tt.want)
		}
	}
}
