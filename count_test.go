package main

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	// Tests are table tests per Go preferred style.
	tests := []struct {
		name  string
		words []string
		want  map[string]uint
	}{
		{
			name:  "empty",
			words: []string{},
			want:  map[string]uint{},
		},
		{
			name:  "one word",
			words: []string{"single"},
			want:  map[string]uint{},
		},
		{
			name:  "one bigram",
			words: []string{"hey", "there"},
			want: map[string]uint{
				"hey there": 1,
			},
		},
		{
			name:  "two bigrams",
			words: []string{"this", "is", "cool"},
			want: map[string]uint{
				"this is": 1,
				"is cool": 1,
			},
		},
		{
			name:  "two bigrams, the same",
			words: []string{"dance", "dance", "dance"},
			want: map[string]uint{
				"dance dance": 2,
			},
		},
		{
			name:  "two bigrams, three times",
			words: []string{"dreaming", "dreaming", "dreaming", "dreaming"},
			want: map[string]uint{
				"dreaming dreaming": 3,
			},
		},
		{
			name:  "some repetitions by two",
			words: []string{"i", "like", "to", "dance", "dance", "dance", "also", "i", "like", "to", "sing"},
			want: map[string]uint{
				"i like":      2,
				"like to":     2,
				"to dance":    1,
				"dance dance": 2,
				"dance also":  1,
				"also i":      1,
				"to sing":     1,
			},
		},
		{
			name:  "repetitions by three",
			words: []string{"bells", "jingle", "bells", "jingle", "bells", "jingle", "all", "the", "way"},
			want: map[string]uint{
				"jingle bells": 2,
				"bells jingle": 3,
				"jingle all":   1,
				"all the":      1,
				"the way":      1,
			},
		},
		{
			name:  "task description test case",
			words: []string{"the", "quick", "brown", "fox", "and", "the", "quick", "blue", "hare"},
			want: map[string]uint{
				"the quick":   2,
				"quick brown": 1,
				"brown fox":   1,
				"fox and":     1,
				"and the":     1,
				"quick blue":  1,
				"blue hare":   1,
			},
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
			if got := Count(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failed to count bigrams.\nGot : %v,\nWant: %v", got, tt.want)
			}
		})
	}

}
