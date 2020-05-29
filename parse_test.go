package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	// Tests are table tests per Go preferred style.
	tests := []struct {
		name string
		text string
		want []string
	}{
		{
			name: "empty input",
			text: "",
			want: []string{},
		},
		{
			name: "whitespace",
			text: " . ,\t :-?... .",
			want: []string{},
		},
		{
			name: "one word no whitespace",
			text: "hello",
			want: []string{"hello"},
		},
		{
			name: "one word some uppercase no whitespace",
			text: "hELLo",
			want: []string{"hello"},
		},
		{
			name: "leading whitespace",
			text: "   hey",
			want: []string{"hey"},
		},
		{
			name: "mixed leading whitespace",
			text: " _:: ., \t wow",
			want: []string{"wow"},
		},
		{
			name: "trailing whitespace",
			text: "Hey   ",
			want: []string{"hey"},
		},
		{
			name: "mixed trailing whitespace",
			text: "Wow!... ., \t \r\n  ",
			want: []string{"wow"},
		},
		{
			name: "two words",
			text: "Hello, world",
			want: []string{"hello", "world"},
		},
		{
			name: "two words with mixed whitespace",
			text: "  \tHello,    world!...\r\n",
			want: []string{"hello", "world"},
		},
		{
			name: "real test case",
			text: "The quick brown fox and the quick blue hare.",
			want: []string{"the", "quick", "brown", "fox", "and", "the", "quick", "blue", "hare"},
		},
		{
			name: "some real text",
			text: "Create an application that can take as input any text file and output a histogram of the bigrams in the text.",
			want: []string{"create", "an", "application", "that", "can", "take", "as", "input", "any", "text", "file", "and", "output", "a", "histogram", "of", "the", "bigrams", "in", "the", "text"},
		},
	}
	// Run multiple test cases in parallel. Use test case names to make it clear which test case has failed.
	// Running test cases in parallel doesn't give us much performance improvement for the parsing logic, just a good pattern to follow.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use got/want idiom per Go testing convention.
			// Want is an expected value, got is the result of the tested function.
			// DeepEqual just compares slices (Go arrays) to make sure they are the same or have any diff
			// (in which case we signal that test case has failed).
			if got := Parse(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failed to parse text: %v.\nGot : %v,\nWant: %v", tt.text, got, tt.want)
			}
		})
	}
}
