package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}
	if len(banner) != 95 {
		return fmt.Errorf("error banner has %d entries", len(banner))
	}
}
