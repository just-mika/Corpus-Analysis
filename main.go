package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	//open file
	csvfile, err := os.Open("fake_tweets.csv")
	if err != nil {
		log.Fatal(err)
	}
	
	df := dataframe.ReadCSV(csvfile)
	sel := df.Select(3)
	fmt.Println(sel)
}
