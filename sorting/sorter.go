package sorting

import (
	"fmt"
	"sort"
)

type Sorter interface {
	Sort([]string)
}

type alphabeticSorter struct {
}

const (
	ALPHABET_TYPE int = iota
)

func NewSorter(sortertype int) (Sorter, error) {
	switch sortertype {
	case ALPHABET_TYPE:
		return alphabeticSorter{}, nil
	default:
		return nil, fmt.Errorf("type unknown.")
	}
}

func (sorter alphabeticSorter) Sort(wl []string) {
	sort.Strings(wl)
}
