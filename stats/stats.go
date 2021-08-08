package stats

import (
	"fmt"
	"wordstats/count"
)

type Stats struct {
	WordChannel <-chan string
	count       int
	WordSlice   []string
}

func (stats *Stats) PrintStats() {
	stats.printCount()
	mostLetter, mostAmount, countMap := count.CountLetters(stats.WordSlice)
	for letter, amount := range countMap {
		fmt.Printf("%s: %d\n", letter, amount)
	}
	fmt.Printf("The letter that occurred the most was %s, which was entered %d times\n", mostLetter, mostAmount)
}

func (stats *Stats) UpdateCount() {
	stats.count = len(stats.WordSlice)
}

func (stats *Stats) printCount() {
	fmt.Printf("Count: %d\n", stats.count)
}

func (stats *Stats) PrintWords() {
	for num, word := range stats.WordSlice {
		fmt.Printf("%d: %s\n", num+1, word)
	}
}
