package main

import "testing"

func TestUnpacking(t *testing.T) {
	var tests = []struct {
		origin string
		want   string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.origin, func(t *testing.T) {
			ans, err := Unpack(tt.origin)
			if err != nil {
				t.Error(err)
			}

			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
