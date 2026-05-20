package main

import "strings"

func TrimArtRows(rows []string) []string {
	result := make([]string, len(rows))
	for i, r := range rows {
		result[i] = strings.TrimRight(r, " ")
	}
	return result
}
