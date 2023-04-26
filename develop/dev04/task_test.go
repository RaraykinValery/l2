package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	input := []string{"пятка", "пятак", "тяпка", "Слиток", "листок", "столик", "мышь"}
	want := map[string][]string{"пятка": {"пятак", "пятка", "тяпка"}, "слиток": {"листок", "слиток", "столик"}}

	output := FindAnagramSets(&input)

	for key, expectedAnagrams := range want {
		outputAnagrams, ok := output[key]
		if !ok {
			t.Errorf("Expected key %s not found in output map", key)
			continue
		}

		if !reflect.DeepEqual(expectedAnagrams, outputAnagrams) {
			t.Errorf("Expected value %v for key %s, but got %v", expectedAnagrams, key, outputAnagrams)
		}
	}
}
