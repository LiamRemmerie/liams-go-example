package count

import (
	"fmt"
	"strings"
)

func CountLetters(words []string) (mostLetter string, mostAmount int, countMap map[string]int) {
	countMap = make(map[string]int)
	var b strings.Builder
	for _, word := range words {
		fmt.Fprintf(&b, "%s", strings.ToUpper(word))
	}
	s := b.String()
	upperCaseA := int('A')
	for asciinum := upperCaseA; asciinum < upperCaseA+26; asciinum++ {
		letter := string(asciinum)
		amount := strings.Count(s, letter)
		countMap[letter] = amount
		if amount > mostAmount {
			mostLetter = letter
			mostAmount = amount
		}
	}
	return
}
