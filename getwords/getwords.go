package getwords

import (
	"bufio"
	"context"
	"os"
)

func GetWords(ctx context.Context) <-chan string {
	words := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			case words <- scanner.Text():
			}
		}
		close(words)
	}()

	return words
}
