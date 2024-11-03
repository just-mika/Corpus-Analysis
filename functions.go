package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"os"
	"os/exec"
	"slices"
	"sort"
	"strings"
	"unicode"
)

var stopWords = []string{
	"me", "my", "myself", "we", "our", "ours", "ourselves", "you", "your", "yours", "yourself",
	"yourselves", "he", "him", "his", "himself", "she", "her", "hers", "herself", "it", "its",
	"itself", "they", "them", "their", "theirs", "themselves", "what", "which", "who", "whom",
	"this", "that", "these", "those", "am", "is", "are", "was", "were", "be", "been", "being",
	"have", "has", "had", "having", "do", "does", "did", "doing", "an", "the", "and", "but",
	"if", "or", "because", "as", "until", "while", "of", "at", "by", "for", "with", "about", "against",
	"between", "into", "through", "during", "before", "after", "above", "below", "to", "from", "up",
	"down", "in", "out", "on", "off", "over", "under", "again", "further", "then", "once", "here",
	"there", "when", "where", "why", "how", "all", "any", "both", "each", "few", "more", "most", "other",
	"some", "such", "no", "nor", "not", "only", "own", "same", "so", "than", "too", "very",
	"can", "will", "just", "don", "should", "now"}

func countChar(text string) {
	text = strings.ToLower(text)
	for _, char := range text {
		characters[char]++
	}
}

func detectWord(text string) {
	word := ""
	text = strings.ToLower(text)
	for _, char := range text {
		if unicode.IsLetter(char) {
			word += string(char)
		}
	}
	if len(word) != 0 {
		/*isStopWord := func() bool {
			for _, stopWord := range stopWords {
				if stopWord == word {
					fmt.Println(stopWord + " detected")
					return true
				}
			}
			return false
		}
		if !isStopWord() {*/
		words[word]++
		//}
	}
}

func count(curr []string) {
	for _, word := range curr {
		detectWord(word)
		countChar(word)
	}
}

func printChars() {
	fmt.Println("\nChars")
	for char, count := range characters {
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			fmt.Print(string(char) + " = ")
			fmt.Printf("%d\n", count)
		}
	}
}

func getWordCount() {
	wordCount := 0
	for _, count := range words {
		wordCount += count
	}
	fmt.Println("Word Count: ", wordCount)
}

func uniqueWordCount() {
	wordCount := len(words)
	/*stopWordCount := 0
	for word, _ := range words {
		if slices.Contains(stopWords, word) {
			stopWordCount++
		}
	}
	wordCount -= stopWordCount*/
	fmt.Println("Unique Word Count: ", wordCount)
}

func sortFrequencies[K comparable, V int](list map[K]V) []K {
	keys := make([]K, 0, len(list))

	// iterate over the map and append all keys to our
	// string array of keys
	for key := range list {
		keys = append(keys, key)
	}

	// use the sort method to sort our keys array
	sort.SliceStable(keys, func(i, j int) bool {
		return list[keys[i]] > list[keys[j]]
	})
	return keys
}

func viewList[K comparable, V int](list []K, dict map[K]V) {
	for _, item := range list {
		var str string
		switch i := any(item).(type) {
		case rune:
			str = string(i)
		case string:
			str = i
		}
		fmt.Println(str, "=", dict[item])
	}
}

func getTop20(sortedWords []string) {
	fmt.Println("Top 20 Words: ")
	for i := 0; i < 20; i++ {
		fmt.Println(i+1, ": "+sortedWords[i]+" = ", words[sortedWords[i]])
	}
}

func getStopWords(sortedWords []string) {
	fmt.Println("Stop Words: ")
	wordNum := len(sortedWords)
	stopCount := 0
	for i := 0; i < wordNum && stopCount < 10; i++ {
		if slices.Contains(stopWords, sortedWords[i]) {
			fmt.Println(stopCount+1, ": "+sortedWords[i]+" = ", words[sortedWords[i]])
			stopCount++
		}
	}
}

func wordCloud(sortedWords []string) {

}

func symbolPie() {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Symbol Distribution"}))
	getSymbols := func() []opts.PieData {
		symbols := make([]opts.PieData, 0)
		var str_char string
		for char, count := range characters {
			if unicode.IsSymbol(char) || unicode.IsPunct(char) {
				str_char = string(char)
				symbols = append(symbols, opts.PieData{Name: str_char, Value: count})
			}
		}
		return symbols
	}

	pie.AddSeries("pie", getSymbols()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
		)
	f, err := os.Create("symb_pie.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pie.Render(f)

	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", "symb_pie.html")
	cmd.Start()
}
