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
var postsPerMonth = make(map[string]int)

func main() {
	var file string
	fmt.Print("Enter file name: ")
	fmt.Scan(&file)

	//	Task 1: Load Data File
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
		addMonth(text[2])
		count(current)
	}

	/*	Task 2: Corpus Analysis */
	/*	Word Count: Word count: Calculate the total number of words.*/
	getWordCount()

	/*	Vocabulary Size: Calculate the number of unique words*/
	uniqueWordCount()
	/*	Word Frequency: Calculate the number of times each word appears in the
		corpus, then sort them based on their frequency.*/
	sortedWords := sortFrequencies(words)
	//viewList(sortedWords, words)

	/*	Character Frequency: Calculate the number of times each character appears in
		the corpus (including letters, digits, and symbols), and then sort them based on
		frequency */
	sortedChars := sortFrequencies(characters)
	viewList(sortedChars, characters)
	/*	Frequency analysis. Identify the top 20 most frequent words and their counts.*/
	getTop20(sortedWords)

	/*	Stop Word Identification: Identify 10 common words (e.g., "the," "and," "a") that
		don't provide significant meaning. */
	getStopWords(sortedWords)

	/* Task 3. Data Visualization */
	//A word cloud of the top 20 most frequent words
	//A bar chart / histogram showing the total number of posts per month
	// A pie chart showing the distribution of the different symbols found in the corpus.
	makeCharts(sortedWords)
}
