package csvh

import (
	"testing"
	"os"
	"fmt"
)

func TestColumnize(t *testing.T) {
	data := [][]string {
		[]string{"apples", "bananas", "pears",},
		[]string{"3", "5", "none",},
		[]string{"100", "1000", "1000000000",},
	}
	Columnize(os.Stdout, data, 8, AlnLeft)
	fmt.Printf("\n")
	Columnize(os.Stdout, data, 8, AlnRight)
}
