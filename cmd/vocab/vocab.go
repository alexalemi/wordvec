package main

import (
	"flag"
	"os"

	"github.com/alexalemi/wordvec"
)

var (
	vocabFile string
	minCount  uint64
	output    string
)

func main() {
	flag.StringVar(&vocabFile, "corpus", "", "Corpus to generate vocab from")
	flag.Uint64Var(&minCount, "min-count", 5, "Minimum count to keep")
	flag.StringVar(&output, "output", "", "Output file")
	flag.Parse()

	if vocabFile == "" || output == "" {
		flag.Usage()
		os.Exit(1)
	}

	v := wordvec.NewVocab(vocabFile, minCount)
	v.Save(output)
}
