package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}
	if len(banner) != 95 {
		return fmt.Errorf("error banner has %d entries", len(banner))
	}
	// for r := rune(32); r <= 126; r++ {
	// 	lines, ok := banner[r]
	// 	if !ok {
	// 		return fmt.Errorf("key is absent")
	// 	}
	// 	if len(lines) != 8 {
	// 		return fmt.Errorf("value is not up to 8")
	// 	}
	// }
	for k, v := range banner {
		if k < rune(32) || k > rune(126) {
			return fmt.Errorf("key is absent")
		}
	}
	return nil
}
