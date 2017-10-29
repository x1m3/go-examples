package sorting_test

import (
	"sort"
	"fmt"
	"strings"
)

type ListOfWords struct {
	words []string
}

func NewListOfWords() *ListOfWords {
	return &ListOfWords{words: make([]string, 0)}

}

func (l *ListOfWords) Add(s string) {
	l.words = append(l.words, s)
}

func (l *ListOfWords) String() string {
	return strings.Join(l.words, ", ")

}

// To be sortable, we must implement sort.interface
// It means that we need to create 3 methods:
// Len() int
// Less(i, j int) bool
// Swap(i, j int)

func (l *ListOfWords) Len() int {
	return len(l.words)
}

// CAUTION. i and j are indexes
func (l *ListOfWords) Less(i, j int) bool {
	return strings.Compare(l.words[i], l.words[j]) < 0
}

// Swap is very funny in go
func (l *ListOfWords) Swap(i, j int) {
	l.words[i], l.words[j] = l.words[j], l.words[i]
}

func ExampleSortObject() {

	l := NewListOfWords()
	l.Add("Zurich")
	l.Add("Barcelona")
	l.Add("Paris")
	l.Add("London")
	l.Add("Geneve")
	l.Add("Munich")
	l.Add("Roma")

	sort.Sort(l)
	fmt.Println(l)

	// Output: Barcelona, Geneve, London, Munich, Paris, Roma, Zurich

}
