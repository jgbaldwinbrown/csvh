package csvh

import (
	"fmt"
	"io"
)

func RepeatString(w io.Writer, str string, reps int) (n int, err error) {
	for i := 0; i < reps; i++ {
		npr, e := fmt.Fprintf(w, "%s", str)
		n += npr
		if e != nil {
			return n, e
		}
	}
	return n, nil
}

type Alignment int

const (
	AlnRight Alignment = iota
	AlnLeft
)

func Columnize(w io.Writer, table [][]string, space int, aln Alignment) (n int, err error) {
	if len(table) < 1 {
		return 0, nil
	}
	lens := make([]int, 0, len(table[0]))
	for _, col := range table[0] {
		lens = append(lens, len(col))
	}
	for _, row := range table[1:] {
		for i, col := range row {
			if len(col) > lens[i] {
				lens[i] = len(col)
			}
		}
	}

	for _, row := range table {
		for i, col := range row {
			if i > 0 {
				npr, e := RepeatString(w, " ", space)
				n += npr
				if e != nil {
					return n, e
				}
			}
			if aln == AlnLeft {
				npr, e := fmt.Fprintf(w, "%s% *s", col, lens[i] - len(col), "")
				n += npr
				if e != nil {
					return n, e
				}
			} else {
				npr, e := fmt.Fprintf(w, "% *s", lens[i], col)
				n += npr
				if e != nil {
					return n, e
				}
			}
		}
		npr, e := fmt.Fprintf(w, "\n")
		n += npr
		if e != nil {
			return n, e
		}
	}
	return n, nil
}
