package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
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

func count(curr []string) {
	for _, word := range curr {
		if len(word) > 1 {
			word = strings.ToLower(word)
			words[word]++
		}
		countChar(word)
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

func addMonth(date string) {
	month := date[5:7]

	switch month {
	case "01":
		postsPerMonth["Jan"]++
	case "02":
		postsPerMonth["Feb"]++
	case "03":
		postsPerMonth["Mar"]++
	case "04":
		postsPerMonth["Apr"]++
	case "05":
		postsPerMonth["May"]++
	case "06":
		postsPerMonth["Jun"]++
	case "07":
		postsPerMonth["Jul"]++
	case "08":
		postsPerMonth["Aug"]++
	case "09":
		postsPerMonth["Sep"]++
	case "10":
		postsPerMonth["Oct"]++
	case "11":
		postsPerMonth["Nov"]++
	case "12":
		postsPerMonth["Dec"]++
	default:
		fmt.Println("Invalid month")
	}
}

func barGraph() *charts.Bar {
	var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	getData := func() []opts.BarData {
		items := make([]opts.BarData, 0)
		for _, month := range months {
			count := postsPerMonth[month]
			items = append(items, opts.BarData{Name: month, Value: count})
		}
		return items
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Posts per Month",
		}))

	bar.SetXAxis(months).
		AddSeries("Bar Graph", getData())
	return bar
}

func wordCloud(sortedWords []string) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Top 20 words",
		}))
	getWords := func() []opts.WordCloudData {
		wcData := make([]opts.WordCloudData, 0)
		for i := 0; i < 20; i++ {
			wcData = append(wcData, opts.WordCloudData{Name: sortedWords[i], Value: words[sortedWords[i]]})
		}
		return wcData
	}
	wc.AddSeries("Word Cloud", getWords()).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "cardioid",
				}),
		)
	return wc
}

func symbolPie() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Symbols"}))
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

	pie.AddSeries("Symbol Pie", getSymbols()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
		)
	return pie
}

func makeCharts(sortedWords []string) {
	page := components.NewPage()
	page.AddCharts(
		wordCloud(sortedWords),
		barGraph(),
		symbolPie(),
	)
	f, err := os.Create("charts.html")
	if err != nil {
		log.Fatal(err)
	}

	page.Render(io.MultiWriter(f))
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", "charts.html")
	cmd.Start()
}
