package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"wordstats/getwords"
	"wordstats/stats"
)

const forbidden string = ".?,!:; "

func main() {
	stats := new(stats.Stats)
	fmt.Println("Hello there! Do you have some words for me? Please separate them with a newline!\nPress Ctrl+C when you're done.")
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	stats.WordChannel = getwords.GetWords(ctx)
	for word := range stats.WordChannel {
		word, err := cleanInput(word)
		if err != nil {
			panic(err)
		}
		stats.WordSlice = append(stats.WordSlice, word)
	}
	stats.UpdateCount()
	stats.PrintWords()
	stats.PrintStats()
}

func cleanInput(w string) (out string, err error) {
	w = strings.TrimSpace(w)
	if strings.ContainsAny(w, forbidden) {
		err = fmt.Errorf("please enter one word, not a sentence.")
	} else if strings.Contains(w, "'") {
		err = fmt.Errorf("sorry, I don't like contractions.")
	} else {
		out = w
	}
	return
}
