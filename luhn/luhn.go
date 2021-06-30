package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	trim := strings.ReplaceAll(s, " ", "")
	sum := 0
	doubleIndex := len(trim) - 1

	for i := len(trim); i > 0; i-- {
		b := trim[i-1]
		r := rune(b)
		if !unicode.IsNumber(r) {
			return false
		}
		toInt, _ := strconv.Atoi(string(b))

		if i == doubleIndex {
			doubleIndex = i - 1
			toInt = toInt * 2
			if toInt > 9 {
				toInt = toInt - 9
			}
		}
		sum += toInt
	}
	return sum%10 == 0
}
