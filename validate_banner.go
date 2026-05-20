package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}
}
