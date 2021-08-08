package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"wordstats/getwords"
	"wordstats/stats"
)

func main() {
	stats := new(stats.Stats)
	fmt.Println("Hello there! Do you have some words for me? Please separate them with a newline!\nPress Ctrl+C when you're done.")
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	stats.WordChannel = getwords.GetWords(ctx)
	for word := range stats.WordChannel {
		stats.WordSlice = append(stats.WordSlice, word)
	}
	stats.UpdateCount()
	stats.PrintWords()
	stats.PrintStats()
}
