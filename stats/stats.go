package stats

import (
	"fmt"
)

type Stats struct {
	WordChannel <-chan string
	count       int
	WordSlice   []string
}

func (stats *Stats) PrintStats() {
	stats.printCount()
}

func (stats *Stats) UpdateCount() {
	stats.count = len(stats.WordSlice)
}

func (stats *Stats) printCount() {
	fmt.Printf("%d", stats.count)
}

func (stats *Stats) PrintWords() {
	for num, word := range stats.WordSlice {
		fmt.Printf("%d: %s\n", num+1, word)
	}
}
