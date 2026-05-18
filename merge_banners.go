package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	m := make(map[rune][]string)
	for k, v := range base {
		m[k] = v
	}
	for k, v := range priority {
		m[k] = v
	}
	return m
}
