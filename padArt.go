package main

import (
	"strings"
)

func PadArtRows(rows []string, width int) []string {
	var sli = []string{}
	for _, r := range rows {
		padding := width - len(r)
		if padding < 0 {
			sli = append(sli, r)
		} else {
			sli = append(sli, r+strings.Repeat(" ", padding))

		}
	}
	return sli
}
