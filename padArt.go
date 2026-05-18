package main

import (
	"fmt"
	"strings"
)

func PadArtRows(rows []string, width int) ([]string, int) {
	for i := 0; i < len(rows); i++ {
		padding := width - len(rows[i])
		if padding < 0 {
			continue
		}
		rows[i] += strings.Repeat(" ", padding)
	}
	return rows, len(rows[0])
}
func main() {
	fmt.Println(PadArtRows([]string{"hello world"}, 8))
}
