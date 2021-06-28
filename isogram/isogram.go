package isogram

import "strings"

func IsIsogram(word string) bool {
	if word == "" {
		return true
	}

	letters := make(map[string]int)
	trim := strings.ReplaceAll(word, " ", "")
	replace := strings.ReplaceAll(trim, "-", "")
	lower := strings.ToLower(replace)
	isIsogram := true

	for _, c := range lower {
		if letters[string(c)] == 1 {
			isIsogram = false
			break
		}
		letters[string(c)] += 1
	}

	return isIsogram
}
