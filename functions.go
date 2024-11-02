package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"unicode"
)

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
		words[word]++
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

func printWords() []string {
	fmt.Println("Words")

	// make an array of type string to store our keys
	keys := []string{}

	// iterate over the map and append all keys to our
	// string array of keys
	for key := range words {
		keys = append(keys, key)
	}

	// use the sort method to sort our keys array
	sort.SliceStable(keys, func(i, j int) bool {
		return words[keys[i]] > words[keys[j]]
	})
	//i := 0
	for _, key := range keys {
		//if i != 20 {
		fmt.Print(key + " = ")
		fmt.Printf("%d\n", words[key])
		//i++
		//} else {
		//	break
		//}
	}
	return keys
}

func wordCloud(sortedWords []string) {

}

func symbolPie() {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Symbol Distribution"}))
	data := func() []opts.PieData {
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

	pie.AddSeries("pie", data()).
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
