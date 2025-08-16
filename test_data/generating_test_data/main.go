package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func randomTitle(wordCount uint) string {
	words := make([]string, wordCount+1)

	for i := range wordCount {
		words[i] = gofakeit.NounCommon()
	}

	words[len(words)-1] = uuid.NewString()
	return strings.Join(words, "/")
}

func requirementGenerator() []string {
	titleWordCount := 10
	title := gofakeit.LoremIpsumSentence(titleWordCount)
	description := gofakeit.LoremIpsumParagraph(2, 3, 20, "\n")
	path := randomTitle(4)

	return []string{
		uuid.NewString(),
		title,
		path,
		description,
	}
}

func main() {
	lineCount := 5000
	records := make([][]string, lineCount)

	for lineNumber := range lineCount {
		records[lineNumber] = requirementGenerator()
	}

	f, err := os.Create("requirements.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
