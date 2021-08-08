package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"wordstats/getwords"
)

func main() {
	fmt.Println("Hello there! Do you have some words for me? Please separate them with a newline!\nPress Ctrl+C when you're done.")
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	words := getwords.GetWords(ctx)
	wordslice := make([]string, 0)
	for word := range words {
		wordslice = append(wordslice, word)
	}
	for num, word := range wordslice {
		fmt.Printf("%d: %s\n", num+1, word)
	}
}
