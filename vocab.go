package wordvec

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Word is a struct holding the count of the word and the word itself
type Word struct {
	Count uint64
	Word  string
}

// WordList is a list of words, and implements the sort interface
type WordList []Word

// Vocab is a list of words and a lookup table
type Vocab struct {
	Words  WordList
	Lookup map[string]int
}

func (w WordList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w WordList) Len() int {
	return len(w)
}
func (w WordList) Less(i, j int) bool {
	if w[i].Count == w[j].Count {
		return w[i].Word < w[j].Word
	}
	return w[i].Count > w[j].Count
}

// mapToList converts a map to a WordList
func mapToList(m map[string]uint64) WordList {
	p := make(WordList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Word{Word: k, Count: v}
		i++
	}
	sort.Sort(p)
	return p
}

func (v Vocab) Len() int {
	return len(v.Words)
}

// Pk gives the location of a word
func (v Vocab) Pk(word string) (int, bool) {
	pk, ok := v.Lookup[word]
	return pk, ok
}

// Add adds one to the count of a word, or adds a new word to the vocab
func (v *Vocab) Add(word string) {
	if pk, ok := v.Pk(word); ok {
		v.Words[pk].Count++
	} else {
		v.Lookup[word] = len(v.Words)
		v.Words = append(v.Words, Word{Word: word, Count: 1})
	}
}

// New creates a new sorted wordlist from the given corpus
func NewVocab(corpus string, minCount uint64) *Vocab {
	log.Println("Reading new corpus from", corpus)
	file, err := os.Open(corpus)
	if err != nil {
		log.Fatal("Cannot open corpus", err)
	}
	defer file.Close()

	counts := make(map[string]uint64)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	i := 0
	for scanner.Scan() {
		word := scanner.Text()
		if c, ok := counts[word]; ok {
			counts[word] = c + 1
		} else {
			counts[word] = 1
		}
		i++
		if i%100000 == 0 {
			fmt.Printf("Processed %vk tokens\r", i/100000)
		}
	}
	fmt.Printf("\n")
	log.Printf("Finished. Found %v words\n", len(counts))

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning corpus", err)
	}

	// purge low counts
	for k, v := range counts {
		if v < minCount {
			delete(counts, k)
		}
	}
	log.Printf("%v words made the cut\n", len(counts))

	wordList := mapToList(counts)
	v := Vocab{Words: wordList}
	v.Lookup = make(map[string]int)
	for i, w := range wordList {
		v.Lookup[w.Word] = i
	}
	return &v
}

func LoadVocab(flname string) *Vocab {
	log.Println("Loading vocab from", flname)
	file, err := os.Open(flname)
	if err != nil {
		log.Fatal("Cannot open vocab file", err)
	}
	defer file.Close()

	v := Vocab{}
	v.Lookup = make(map[string]int)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		if len(parts) < 2 {
			log.Fatal("Line doesn't contain counts", line)
		}

		word := parts[0]
		count, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			log.Fatal("Error parsing count in line ", err)
		}
		v.Words = append(v.Words, Word{Word: word, Count: count})
		v.Lookup[word] = i
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning vocab", err)
	}

	return &v
}

func (v Vocab) Save(flname string) {
	log.Println("Saving vocab to", flname)
	file, err := os.Create(flname)
	if err != nil {
		log.Fatal("Cannot open file for saving", err)
	}
	defer file.Close()

	for _, w := range v.Words {
		fmt.Fprintf(file, "%v\t%v\n", w.Word, w.Count)
	}
}
