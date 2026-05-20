package main

func PadArtRows(rows []string, width int) []string {
	sli := []string{}
	for _, r := range rows {
		padding := width - len(r)
		if padding <= 0 {
			sli = append(sli, r)
		}
	}
}
