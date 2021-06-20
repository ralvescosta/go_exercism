package acronym

import (
	"strings"
	"unicode"
)

func splitDigit(s string) []string {
	for _, c := range s {
		if c == '-' {
			return strings.Split(s, "-")
		}
	}
	return []string{s}
}

func Abbreviate(s string) string {
	spited := strings.Split(s, " ")
	abbreviate := ""
	for _, item := range spited {
		spitedInDigit := splitDigit(item)
		for _, item2 := range spitedInDigit {
			for _, r := range item2 {
				if unicode.IsLetter(r) {
					abbreviate += strings.ToUpper(string(r))
					break
				}
			}
		}
	}

	return abbreviate
}
