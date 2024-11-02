package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var words = make(map[string]int)
var characters = make(map[rune]int)

func main() {
	var file string
	fmt.Print("Enter file name: ")
	fmt.Scan(&file)

	//Task 1: Load Data File
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	for {
		text, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		current := strings.Fields(text[3])
		count(current)
	}

	/* Task 2: Corpus Analysis
	Word Count: Word count: Calculate the total number of words.

	Vocabulary Size: Calculate the number of unique words

	Word Frequency: Calculate the number of times each word appears in the
	corpus, then sort them based on their frequency.

	Character Frequency: Calculate the number of times each character appears in
	the corpus (including letters, digits, and symbols), and then sort them based on
	frequency

	Frequency analysis. Identify the top 20 most frequent words and their counts.

	Stop Word Identification: Identify 10 common words (e.g., "the," "and," "a") that
	don't provide significant meaning.
	*/
	sorted := printWords()
	if len(characters) > 0 {
		symbolPie()
	}

	wordCloud(sorted)
	//printChars()
}
