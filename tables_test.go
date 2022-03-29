package main

import (
	"fmt"
	"testing"
)

func TestWordWrap(t *testing.T) {
	var tests = []struct {
		word    string
		colSize int
		want    string
	}{
		{"enslavement of beauty", 8, "enslave…"},
		{"sangue de bode", 10, "sangue de…"},
		{"first fragment", 7, "first …"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%d", tt.word, tt.colSize)
		t.Run(testname, func(t *testing.T) {
			ans := wordWrap(tt.word, tt.colSize)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestCheckLimit(t *testing.T) {
	var tests = []struct {
		limit string
		want  string
	}{
		{"20", "20"},
		{"A", "10"},
		{"8", "8"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.limit)
		t.Run(testname, func(t *testing.T) {
			ans := checkLimit(tt.limit)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
