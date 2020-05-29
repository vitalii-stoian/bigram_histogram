# Bigram histogram

## Objective

Take as input any text file and output a histogram of the bigrams in the text.

## Definitions

A `bigram` is any two adjacent words in the text disregarding case and punctuation. 
A `histogram` is the count of how many times that particular `bigram` occurred in the text.

## Example:

Given the text: `The quick brown fox and the quick blue hare.` The bigrams with their counts
would be:
* `the quick` : 2
* `quick brown` : 1
* `brown fox` : 1
* `fox and` : 1
* `and the` : 1
* `quick blue` : 1
* `blue hare` : 1

## Solution components

The problem consists of three parts:
1. Read a contents of the input file, which name is passed as a command line parameter.
1. Parse the contents of the file to get words. 
1. Count the histogram distribution of the bigrams.

## Assumptions

* Valid word characters are:
  * lowercase letters `a-z`
  * uppercase letters `A-Z`
  * digits `0-9`
* Words are converted into lowercase as a result of parsing the input text
* Number of individual bigram occurrences is less than `2^32-1` so it fits into standard numeric type `uint` (unsigned integer). Note: on some platforms it could fit `2^64-1`, however, we assume the max is `2^32-1`.

## Solution details

### Parsing

* `isLetter` lambda function is used to validate each character to determine if it's an allowed character for a word.
* Leading non-word characters are skipped until first valid character is found.
* `wordStartPos` is used to keep track of the starting position of each new word.
* Parsing consists of a single loop through all input text.
  * Valid letters are skipped.
  * New word is appended to a slice (array) of words once a non-word character is found.
  * Non-word characters are skipped to move to the beginning of the next word.
* Finally, the last word is appended to a slice of words in case the text ends with a word character, so the loop body doesn't handle this corner case.
* A slice of words is returned as a result of `Parse()`.

Runtime complexity of this part is linear: `O(N)`, where `N` is a total number of characters in the input text. There's no way to make it better because we cannot skip any input characters not to miss any word.

### Counting

* Bigram is a neighbor pair of parsed words, which leads to a straightforward solution to go over all parsed words and count every next bigram.
* Hashmap is used to keep track of bigrams and their occurences.
  * Go `map` is implemented as a hash container under the hood.
  * Insert, update, remove operations are amortized constant time `O(1)`.
  * In the worst case, all operation might take `O(K)` where `K` is a number of elements in the hashmap. This happens in case of collisions when has function leads different keys to the same bucket, causing the values to be stored in a linked list, resulting in te linear runtime complexity for all operations. However, string keys in real life are distributed pretty uniformly, so we can assume the probability of collision is rather low.
* Each bigram is either inserted into a hasmap (if didn't exist) with value `1` or is accessed (if existed) and its value is incremented.
* Populated in this way hashmap represents a target histogram distribution, so it's returned as a result of `Count()`.

Runtime complexity of this part is `O(M)`, where `M` is a total number of unique (distinct) bigrams. Hashmap grows to `M` elements. Each operation is `O(1)`.

### Overall runtime complexity of the solution

`O(N) + O(M) = O(N)`, where `N` is the number of characters in the input text and `M` is the number of bigrams. Obviously, `M < N`, therefore `O(N)` dominates over `O(M)`.

### Potential optimization

Overall runtime complexity could be just `O(N)`, not `O(N) + O(M)`. This could be achieved if both parsing and bigram countimg happens at the same time during a single pass of the input text.

However, this approach is not chosen because it would have made the solution harder to understand and harder to test. Two separate steps are more readable and could be tested with their own set of test cases, which are the main reasons to not prematurely optimize the solution. On the other hand, the solution becomes optimal if this optimization is implemented.

## Testing

Both parsing and counting are covered with unit tests. Go style table tests are used, so each of two test functions contain all related test cases in a slice of structs, which is iterated and checked.

## Output

```shell
 $ pwd                                                                                                          
/home/vitalii/go/src/github.com/vitalii-stoian/bigram_histogram      

 $ go build

 $ ll bigram_histogram
2.4M -rwxr-xr-x 1 vitalii vitalii 2.4M May 28 23:07 bigram_histogram

 $ cat input
The quick brown fox and the quick blue hare.

 $ ./bigram_histogram input
2020/05/28 23:07:54 Bigrams histogram:
{
  "and the": 1,
  "blue hare": 1,
  "brown fox": 1,
  "fox and": 1,
  "quick blue": 1,
  "quick brown": 1,
  "the quick": 2
}

 $ go test
PASS
ok      github.com/vitalii-stoian/bigram_histogram      0.002s

```