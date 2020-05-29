package main

import "strings"

// Count takes a slice of words and returns a histogram of the bigrams.
//
// Complexity: O(N) where N is the number of words (each individual bigram contributes just an amortized O(1) time because of using a hashmap).
//
// It's assumed that the number of individual bigram occurrences is less than the max value of int type,
// so we don't do an extra check to handle that case.
//
// It could have been possible to count bigrams while we parse them from the imput text.
// This way, a single pass would have been required unlike two passes when we first parse,
// and the count. However, splitting parsing and counting makes much more sense in terms of
// better readability and testability that saving a bit of performance.
// However, this might be important in production when we need to super optimize and have already optimized all other places :)
func Count(words []string) map[string]uint {
	// Go map is a hasmap under the hood.
	// Complexity of insert and update operations are amortized O(1).
	counts := map[string]uint{}

	// Go from zero to one less than the last word to process a pair of adjacent words on each iteration.
	for i := 0; i < len(words)-1; i++ {
		// Each bigram is two adjacent words joined with a single space.
		// Rely on the assumption that words already contain only lower case letters (done at parsing time).
		bigram := strings.Join([]string{words[i], words[i+1]}, " ")

		// Simply increment the number of occurrences of the current bigram.
		// If there were no bigram in the hasmap, the entry will be created with the bigram as the key and the value 1.
		// If there was a bigram as a key in the hasmap already, it's value will be incremented.
		// This operation takes amortized constant time
		// (bucket is determined based on a hash of the bigram, collisions are rare for strings in practice).
		counts[bigram]++
	}

	return counts
}
